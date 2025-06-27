package actor

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/browser"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/captcha"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/httpserver"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/proxies"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/router"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/storage"

	"github.com/tidwall/gjson"
	"reflect"
)

type Actor struct {
	Browser     *browser.Browser
	Proxy       *proxies.Proxy
	Captcha     *captcha.Captcha
	storage     *storage.Storage
	Server      *httpserver.Server
	Router      *router.Router
	closeFun    []func() error
	datasetId   string
	namespaceId string
	bucketId    string
	queueId     string
	vectorId    string
}

const (
	typeGrpc = "grpc"
	typeHttp = "http"
)

// New creates a new Actor.
func New() *Actor {
	var actor = new(Actor)
	actor.storage = storage.NewStorage(typeHttp)
	actor.Browser = browser.NewBrowser(typeHttp)
	actor.Captcha = captcha.NewCaptcha(typeHttp)
	actor.Proxy = proxies.NewProxy(typeHttp)
	actor.Router = router.New(typeHttp)
	actor.Server = httpserver.New()

	actor.datasetId = env.Env.Actor.DatasetId
	actor.namespaceId = env.Env.Actor.KvNamespaceId
	actor.bucketId = env.Env.Actor.BucketId
	actor.queueId = env.Env.Actor.QueueId
	actor.vectorId = env.Env.Actor.VectorId
	return actor
}

// Close closes the actor.
func (a *Actor) Close() {
	for _, f := range a.closeFun {
		_ = f()
	}
}

// Input get input data from env.
func (a *Actor) Input(data any) error {
	input, err := a.GetValue(context.Background(), "INPUT")
	if err != nil {
		return err
	}
	input = gjson.Parse(input).String()
	tf := reflect.TypeOf(data)
	if tf.Kind() != reflect.Ptr {
		return errors.New("data must be ptr")
	}
	return json.Unmarshal([]byte(input), data)
}

func (a *Actor) Start() error {
	return a.Server.Start(fmt.Sprintf(":%s", env.Env.Actor.HttpPort))
}

/**
 * KV store convenience methods with environment variables
 */

// ListNamespaces List all available namespaces
func (a *Actor) ListNamespaces(ctx context.Context, page int, pageSize int, desc bool) (*storage.NamespacesResponse, error) {
	return a.storage.KV.ListNamespaces(ctx, int64(page), int64(pageSize), desc)
}

// CreateNamespace Create a new namespace
func (a *Actor) CreateNamespace(ctx context.Context, name string) (namespaceId string, namespaceName string, err error) {
	return a.storage.KV.CreateNamespace(ctx, name)
}

// GetNamespace Get a namespace
func (a *Actor) GetNamespace(ctx context.Context, namespaceName string) (*storage.KvNamespaceItem, error) {
	return a.storage.KV.GetNamespace(ctx, namespaceName)
}

// DelNamespace Delete a namespace
func (a *Actor) DelNamespace(ctx context.Context) (bool, error) {
	return a.storage.KV.DelNamespace(ctx, a.namespaceId)
}

// RenameNamespace Rename a namespace
func (a *Actor) RenameNamespace(ctx context.Context, name string) (ok bool, namespaceName string, err error) {
	return a.storage.KV.RenameNamespace(ctx, a.namespaceId, name)
}

// ListKeys List keys in a namespace
func (a *Actor) ListKeys(ctx context.Context, page int, pageSize int) (*storage.KvKeys, error) {
	return a.storage.KV.ListKeys(ctx, a.namespaceId, int64(page), int64(pageSize))
}

// SetValue Set a key-value pair in the default namespace (from environment variable)
func (a *Actor) SetValue(ctx context.Context, key string, value string, expiration uint) (bool, error) {
	return a.storage.KV.SetValue(ctx, a.namespaceId, key, value, expiration)
}

// DeleteValue Delete a value from a namespace
func (a *Actor) DeleteValue(ctx context.Context, key string) (bool, error) {
	return a.storage.KV.DelValue(ctx, a.namespaceId, key)
}

// BulkSetValue Bulk set multiple key-value pairs in a namespace
func (a *Actor) BulkSetValue(ctx context.Context, data []storage.BulkItem) (successCount int64, err error) {
	return a.storage.KV.BulkSetValue(ctx, a.namespaceId, data)
}

// BulkDelValue Bulk delete multiple keys from a namespace
func (a *Actor) BulkDelValue(ctx context.Context, keys []string) (bool, error) {
	return a.storage.KV.BulkDelValue(ctx, a.namespaceId, keys)
}

// GetValue Get a value by key from the default namespace (from environment variable)
func (a *Actor) GetValue(ctx context.Context, key string) (string, error) {
	return a.storage.KV.GetValue(ctx, a.namespaceId, key)
}

/**
 * Dataset convenience methods
 */

// ListDatasets  list all available datasets
func (a *Actor) ListDatasets(ctx context.Context, page int64, pageSize int64, desc bool) (*storage.ListDatasetsResponse, error) {
	return a.storage.Dataset.ListDatasets(ctx, page, pageSize, desc)
}

// CreateDataset create a new dataset
func (a *Actor) CreateDataset(ctx context.Context, name string) (id string, datasetName string, err error) {
	return a.storage.Dataset.CreateDataset(ctx, name)
}

// UpdateDataset update an existing dataset
func (a *Actor) UpdateDataset(ctx context.Context, name string) (ok bool, datasetName string, err error) {
	return a.storage.Dataset.UpdateDataset(ctx, a.datasetId, name)
}

// DeleteDataset delete a dataset
func (a *Actor) DeleteDataset(ctx context.Context) (bool, error) {
	return a.storage.Dataset.DelDataset(ctx, a.datasetId)
}

// AddItems Add items to the default dataset (from environment variable)
func (a *Actor) AddItems(ctx context.Context, items []map[string]any) (bool, error) {
	return a.storage.Dataset.AddItems(ctx, a.datasetId, items)
}

// GetItems Get items from the default dataset (from environment variable)
func (a *Actor) GetItems(ctx context.Context, page int, pageSize int, desc bool) (*storage.ItemsResponse, error) {
	return a.storage.Dataset.GetItems(ctx, a.datasetId, page, pageSize, desc)
}

/**
 * Queue convenience methods with environment variables
 */

// ListQueues List all available queues
func (a *Actor) ListQueues(ctx context.Context, page int64, pageSize int64, desc bool) (*storage.ListQueuesResponse, error) {
	return a.storage.Queue.ListQueues(ctx, page, pageSize, desc)
}

// CreateQueue Create a new queue
func (a *Actor) CreateQueue(ctx context.Context, req *storage.CreateQueueReq) (queueId string, queueName string, err error) {
	return a.storage.Queue.CreateQueue(ctx, req)
}

// GetQueue Get a queue by name
func (a *Actor) GetQueue(ctx context.Context, name string) (*storage.Item, error) {
	return a.storage.Queue.GetQueue(ctx, a.queueId, name)
}

// UpdateQueue Update a queue
func (a *Actor) UpdateQueue(ctx context.Context, name string, description string) error {
	return a.storage.Queue.UpdateQueue(ctx, a.queueId, name, description)
}

// DeleteQueue Delete a queue
func (a *Actor) DeleteQueue(ctx context.Context) error {
	return a.storage.Queue.DeleteQueue(ctx, a.queueId)
}

// PushMessage Push a message to the default queue (from environment variable)
func (a *Actor) PushMessage(ctx context.Context, req storage.PushQueue) (string, error) {
	return a.storage.Queue.Push(ctx, a.queueId, req)
}

// PullMessage Pull a message from the default queue (from environment variable)
func (a *Actor) PullMessage(ctx context.Context, size int32) (storage.GetMsgResponse, error) {
	return a.storage.Queue.Pull(ctx, a.queueId, size)
}

// AckMessage Acknowledge a message in the default queue (from environment variable)
func (a *Actor) AckMessage(ctx context.Context, msgId string) error {
	return a.storage.Queue.Ack(ctx, a.queueId, msgId)
}

/**
 * Object storage convenience methods with environment variables
 */

// ListBuckets List all available buckets
func (a *Actor) ListBuckets(ctx context.Context, page int, pageSize int) (*storage.ListBucketsResponse, error) {
	return a.storage.Object.ListBuckets(ctx, page, pageSize)
}

// CreateBucket Create a new bucket
func (a *Actor) CreateBucket(ctx context.Context, name string, description string) (bucketId string, bucketName string, err error) {
	return a.storage.Object.CreateBucket(ctx, name, description)
}

// DeleteBucket Delete a bucket
func (a *Actor) DeleteBucket(ctx context.Context) (bool, error) {
	return a.storage.Object.DeleteBucket(ctx, a.bucketId)
}

// GetBucket Get a bucket
func (a *Actor) GetBucket(ctx context.Context) (*storage.Bucket, error) {
	return a.storage.Object.GetBucket(ctx, a.bucketId)
}

// List list objects in a bucket
func (a *Actor) List(ctx context.Context, fuzzyFileName string, page int64, pageSize int64) (*storage.ListObjectsResponse, error) {
	return a.storage.Object.ListObjects(ctx, a.bucketId, fuzzyFileName, page, pageSize)
}

// GetObject Get an object from the default bucket (from environment variable)
func (a *Actor) GetObject(ctx context.Context, objectId string) ([]byte, error) {
	return a.storage.Object.GetObject(ctx, a.bucketId, objectId)
}

// PutObject Upload an object to the default bucket (from environment variable)
func (a *Actor) PutObject(ctx context.Context, filename string, data []byte) (string, error) {
	return a.storage.Object.PutObject(ctx, a.bucketId, filename, data)
}

// DeleteObject Delete an object from a bucket
func (a *Actor) DeleteObject(ctx context.Context, objectId string) (bool, error) {
	return a.storage.Object.DeleteObject(ctx, a.bucketId, objectId)
}

/**
 * Vector storage convenience methods with environment variables
 */

// ListCollections retrieves a list of vector collections with pagination and sorting options.
func (a *Actor) ListCollections(ctx context.Context, page int64, pageSize int64, desc bool) (*storage.ListCollectionsResponse, error) {
	return a.storage.Vector.ListCollections(ctx, page, pageSize, desc)
}

// CreateCollections creates a new vector collection with the provided request parameters.
func (a *Actor) CreateCollections(ctx context.Context, req *storage.CreateCollectionRequest) (*storage.CreateCollectionResponse, error) {
	return a.storage.Vector.CreateCollections(ctx, req)
}

// UpdateCollection updates the collection information with the provided name and description.
func (a *Actor) UpdateCollection(ctx context.Context, collId string, name string, description string) error {
	return a.storage.Vector.UpdateCollection(ctx, collId, name, description)
}

// DelCollection deletes the collection by its ID.
func (a *Actor) DelCollection(ctx context.Context, collId string) error {
	return a.storage.Vector.DelCollection(ctx, collId)
}

// GetCollection retrieves a collection by its ID.
func (a *Actor) GetCollection(ctx context.Context, collId string) (*storage.Collection, error) {
	return a.storage.Vector.GetCollection(ctx, collId)
}

// CreateDocs inserts new documents into the collection.
func (a *Actor) CreateDocs(ctx context.Context, collId string, docs []*storage.BaseDoc) (*storage.DocOpResponse, error) {
	return a.storage.Vector.CreateDocs(ctx, collId, docs)
}

// UpdateDocs updates existing documents in the collection.
func (a *Actor) UpdateDocs(ctx context.Context, collId string, docs []*storage.Doc) (*storage.DocOpResponse, error) {
	return a.storage.Vector.UpdateDocs(ctx, collId, docs)
}

// UpsertDocs inserts or updates documents in the collection.
func (a *Actor) UpsertDocs(ctx context.Context, collId string, docs []*storage.Doc) (*storage.DocOpResponse, error) {
	return a.storage.Vector.UpsertDocs(ctx, collId, docs)
}

// DelDocs deletes documents from the collection by their IDs.
func (a *Actor) DelDocs(ctx context.Context, collId string, ids []string) (*storage.DocOpResponse, error) {
	return a.storage.Vector.DelDocs(ctx, collId, ids)
}

// QueryDocs queries documents in the collection by vector.
func (a *Actor) QueryDocs(ctx context.Context, collId string, query *storage.QueryVectorParam) ([]*storage.Doc, error) {
	return a.storage.Vector.QueryDocs(ctx, collId, query)
}

// QueryDocsByIds queries documents in the collection by their IDs.
func (a *Actor) QueryDocsByIds(ctx context.Context, collId string, ids []string) (map[string]*storage.Doc, error) {
	return a.storage.Vector.QueryDocsByIds(ctx, collId, ids)
}
