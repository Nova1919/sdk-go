package storage

import (
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/storage/dataset"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/storage/kv"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/storage/object"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/storage/queue"
	"sync"
)

type Storage interface {
	GetKv(namespaceId ...string) kv.KV
	GetObject(bucketId ...string) object.Object
	GetQueue(queueId ...string) queue.Queue
	GetDataset(datasetId ...string) dataset.Dataset
	Close() error
}

var (
	once           = &sync.Once{}
	defaultStorage Storage
)

type StorageHttp struct {
	kvs      map[string]kv.KV
	objs     map[string]object.Object
	queues   map[string]queue.Queue
	datasets map[string]dataset.Dataset
	lock     sync.Mutex
}

// GetKv returns a kv instance
func (sh *StorageHttp) GetKv(namespaceId ...string) kv.KV {
	if len(namespaceId) == 0 {
		if _, ok := sh.kvs[Default]; ok {
			return sh.kvs[Default]
		}
		panic("storage not init")
	}
	sh.lock.Lock()
	defer sh.lock.Unlock()
	if _, ok := sh.kvs[namespaceId[0]]; !ok {
		sh.kvs[namespaceId[0]] = kv.NewKVHttp(namespaceId[0])
	}
	return sh.kvs[namespaceId[0]]
}

// GetObject returns a object instance
func (sh *StorageHttp) GetObject(bucketId ...string) object.Object {
	if len(bucketId) == 0 {
		if _, ok := sh.objs[Default]; ok {
			return sh.objs[Default]
		}
		panic("storage not init")
	}
	sh.lock.Lock()
	defer sh.lock.Unlock()
	if _, ok := sh.objs[bucketId[0]]; !ok {
		sh.objs[bucketId[0]] = object.NewObjHttp(bucketId[0])
	}
	return sh.objs[bucketId[0]]
}

// GetQueue returns a queue instance
func (sh *StorageHttp) GetQueue(queueId ...string) queue.Queue {
	if len(queueId) == 0 {
		if _, ok := sh.queues[Default]; ok {
			return sh.queues[Default]
		}
		panic("storage not init")
	}
	sh.lock.Lock()
	defer sh.lock.Unlock()
	if _, ok := sh.queues[queueId[0]]; !ok {
		sh.queues[queueId[0]] = queue.NewQueueHttp(queueId[0])
	}
	return sh.queues[queueId[0]]
}

// GetDataset returns a dataset instance
func (sh *StorageHttp) GetDataset(datasetId ...string) dataset.Dataset {
	if len(datasetId) == 0 {
		if _, ok := sh.datasets[Default]; ok {
			return sh.datasets[Default]
		}
		panic("storage not init")
	}
	sh.lock.Lock()
	defer sh.lock.Unlock()
	if _, ok := sh.datasets[datasetId[0]]; !ok {
		sh.datasets[datasetId[0]] = dataset.NewDSHttp(datasetId[0])
	}
	return sh.datasets[datasetId[0]]
}

const Default = "default"

// NewStorageHttp returns a storage instance with default key.
func NewStorageHttp() Storage {
	once.Do(func() {
		defaultKv := kv.NewKVHttp()
		defaultObj := object.NewObjHttp()
		defaultQueue := queue.NewQueueHttp()
		defaultDataset := dataset.NewDSHttp()
		sh := &StorageHttp{
			map[string]kv.KV{
				Default: defaultKv,
			},
			map[string]object.Object{
				Default: defaultObj,
			},
			map[string]queue.Queue{
				Default: defaultQueue,
			},
			map[string]dataset.Dataset{
				Default: defaultDataset,
			},
			sync.Mutex{},
		}
		defaultStorage = sh
	})
	return defaultStorage
}

func (sh *StorageHttp) Close() error {
	return nil
}
