package storage

import "context"

type Dataset interface {
	ListDatasets(ctx context.Context, page int64, pageSize int64, desc bool) (*ListDatasetsResponse, error)
	CreateDataset(ctx context.Context, name string) (id string, datasetName string, err error)
	UpdateDataset(ctx context.Context, datasetId string, name string) (ok bool, datasetName string, err error)
	DelDataset(ctx context.Context, datasetId string) (bool, error)
	AddItems(ctx context.Context, datasetId string, items []map[string]any) (bool, error)
	GetItems(ctx context.Context, datasetId string, page int, pageSize int, desc bool) (*ItemsResponse, error)

	Close() error
}

type KV interface {
	ListNamespaces(ctx context.Context, page int, pageSize int, desc bool) (*NamespacesResponse, error)
	CreateNamespace(ctx context.Context, name string) (namespaceId string, namespaceName string, err error)
	GetNamespace(ctx context.Context, namespaceName string) (*KvNamespaceItem, error)
	DelNamespace(ctx context.Context, namespaceId string) (bool, error)
	RenameNamespace(ctx context.Context, namespaceId string, name string) (ok bool, namespaceName string, err error)
	ListKeys(ctx context.Context, namespaceId string, page int, pageSize int) (*KvKeys, error)
	SetValue(ctx context.Context, namespaceId string, key string, value string, expiration uint) (bool, error)
	DelValue(ctx context.Context, namespaceId string, key string) (bool, error)
	BulkSetValue(ctx context.Context, namespaceId string, data []BulkItem) (successCount int64, err error)
	BulkDelValue(ctx context.Context, namespaceId string, keys []string) (bool, error)
	GetValue(ctx context.Context, namespaceId string, key string) (string, error)

	Close() error
}

type Queue interface {
	ListQueues(ctx context.Context, page int64, pageSize int64, desc bool) (*ListQueuesResponse, error)
	CreateQueue(ctx context.Context, req *CreateQueueReq) (queueId string, queueName string, err error)
	GetQueue(ctx context.Context, queueId string, name string) (*Item, error)
	UpdateQueue(ctx context.Context, queueId string, name string, description string) error
	DeleteQueue(ctx context.Context, queueId string) error
	PushMsg(ctx context.Context, queueId string, req PushQueue) (string, error)
	PullMsg(ctx context.Context, queueId string, size int32) (GetMsgResponse, error)
	Ack(ctx context.Context, queueId string, msgId string) error
	Close() error
}

type Object interface {
	ListBuckets(ctx context.Context, page int, pageSize int) (*ListBucketsResponse, error)
	CreateBucket(ctx context.Context, name string, description string) (bucketId string, bucketName string, err error)
	DeleteBucket(ctx context.Context, bucketId string) (bool, error)
	GetBucket(ctx context.Context, bucketId string) (*Bucket, error)
	ListObjects(ctx context.Context, bucketId string, fuzzyFileName string, page int64, pageSize int64) (*ListObjectsResponse, error)
	GetObject(ctx context.Context, bucketId string, objectId string) ([]byte, error)
	PutObject(ctx context.Context, bucketId string, filename string, data []byte) (string, error)
	DeleteObject(ctx context.Context, bucketId string, objectId string) (bool, error)
	Close() error
}
