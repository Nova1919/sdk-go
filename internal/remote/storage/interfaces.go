package storage

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage/models"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage/storage_http"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage/storage_memory"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

type Dataset interface {
	ListDatasets(ctx context.Context, req *models.ListDatasetsRequest) (*models.ListDatasetsResponse, error)
	CreateDataset(ctx context.Context, req *models.CreateDatasetRequest) (*models.Dataset, error)
	UpdateDataset(ctx context.Context, datasetID, name string) (bool, error)
	DelDataset(ctx context.Context, datasetID string) (bool, error)
	GetDataset(ctx context.Context, req *models.GetDataset) (*models.DatasetItem, error)
	AddDatasetItem(ctx context.Context, datasetId string, data []map[string]any) (bool, error)
	Close() error
}

type KV interface {
	ListNamespaces(ctx context.Context, page int64, pageSize int64, desc bool) (*models.KvNamespace, error)
	CreateNamespace(ctx context.Context, req *models.CreateKvNamespaceRequest) (string, error)
	GetNamespace(ctx context.Context, namespaceId string) (*models.KvNamespaceItem, error)
	DelNamespace(ctx context.Context, namespaceId string) (bool, error)
	RenameNamespace(ctx context.Context, namespaceId string, name string) (bool, error)
	SetValue(ctx context.Context, req *models.SetValue) (bool, error)
	ListKeys(ctx context.Context, req *models.ListKeyInfo) (*models.KvKeys, error)
	GetValue(ctx context.Context, namespaceId string, key string) (string, error)
	DelValue(ctx context.Context, namespaceId string, key string) (bool, error)
	BulkSetValue(ctx context.Context, req *models.BulkSet) (int64, error)
	BulkDelValue(ctx context.Context, namespaceId string, keys []string) (bool, error)
	Close() error
}

type Queue interface {
	CreateQueue(ctx context.Context, req *models.CreateQueueRequest) (*models.CreateQueueResponse, error)
	GetQueue(ctx context.Context, req *models.GetQueueRequest) (*models.GetQueueResponse, error)
	GetQueues(ctx context.Context, req *models.GetQueuesRequest) (*models.ListQueuesResponse, error)
	UpdateQueue(ctx context.Context, req *models.UpdateQueueRequest) error
	DelQueue(ctx context.Context, req *models.DelQueueRequest) error
	CreateMsg(ctx context.Context, req *models.CreateMsgRequest) (*models.CreateMsgResponse, error)
	GetMsg(ctx context.Context, req *models.GetMsgRequest) (*models.GetMsgResponse, error)
	AckMsg(ctx context.Context, req *models.AckMsgRequest) error
	Close() error
}

type Object interface {
	ListBuckets(ctx context.Context, page, size int) (*models.Object, error)
	CreateBucket(ctx context.Context, req *models.CreateBucketRequest) (string, error)
	DeleteBucket(ctx context.Context, bucketId string) (bool, error)
	GetBucket(ctx context.Context, bucketId string) (*models.Bucket, error)
	ListObjects(ctx context.Context, req *models.ListObjectsRequest) (*models.ObjectList, error)
	GetObject(ctx context.Context, req *models.ObjectRequest) ([]byte, error)
	DeleteObject(ctx context.Context, req *models.ObjectRequest) (bool, error)
	PutObject(ctx context.Context, req *models.PutObjectRequest) (string, error)
	Close() error
}

type Vector interface {
	ListCollections(ctx context.Context, req *models.ListCollectionsRequest) (*models.ListCollectionsResponse, error)
	CreateCollections(ctx context.Context, req *models.CreateCollectionRequest) (*models.CreateCollectionResponse, error)
	UpdateCollection(ctx context.Context, req *models.UpdateCollectionRequest) error
	DelCollection(ctx context.Context, collId string) error
	GetCollection(ctx context.Context, collId string) (*models.Collection, error)
	CreateDocs(ctx context.Context, req *models.CreateDocsRequest) (*models.DocOpResponse, error)
	UpdateDocs(ctx context.Context, req *models.UpdateDocsRequest) (*models.DocOpResponse, error)
	UpsertDocs(ctx context.Context, req *models.UpsertVectorDocsParam) (*models.DocOpResponse, error)
	DelDocs(ctx context.Context, req *models.DeleteDocsRequest) (*models.DocOpResponse, error)
	QueryDocs(ctx context.Context, req *models.QueryVectorRequest) ([]*models.Doc, error)
	QueryDocsByIds(ctx context.Context, req *models.QueryDocsByIdsRequest) (map[string]*models.Doc, error)
	Close() error
}

type Storage interface {
	Dataset
	KV
	Object
	Queue
	Vector
}

var ClientInterface Storage

func NewClient(serverMode, baseUrl string) {
	if !env.Env.IsOnline {
		serverMode = "dev"
	}
	switch serverMode {
	case "grpc":
		log.Info("grpc...")
	case "dev":
		log.Info("dev...")
		storage_memory.Init()
		ClientInterface = storage_memory.Default()
	default:
		storage_http.Init(baseUrl)
		ClientInterface = storage_http.Default()
	}
}
