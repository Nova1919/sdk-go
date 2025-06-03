package storage

import (
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/storage/dataset"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/storage/kv"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/storage/object"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/storage/queue"
)

type Storage struct {
	Kv      kv.KV
	Object  object.Object
	Queue   queue.Queue
	Dataset dataset.Dataset
}

var (
	defaultStorage Storage
)

func init() {
	defaultStorage = Storage{
		Kv:      kv.NewKVHttp(),
		Object:  object.NewObjHttp(),
		Queue:   queue.NewQueueHttp(),
		Dataset: dataset.NewDSHttp(),
	}
}

func NewStorage() Storage {
	return defaultStorage
}

func (s *Storage) Close() error {
	return nil
}
