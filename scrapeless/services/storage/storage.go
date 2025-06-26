package storage

import (
	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage"
	"sync"
)

type Storage struct {
	*Dataset
	*KV
	*Object
	*Queue
	*Vector
}

var (
	defaultStorage *Storage
	once           = sync.Once{}
)

func init() {
	defaultStorage = &Storage{
		Dataset: &Dataset{},
		KV:      &KV{},
		Object:  &Object{},
		Queue:   &Queue{},
		Vector:  &Vector{},
	}
}

func NewStorage(serverMode string) *Storage {
	once.Do(func() {
		storage.NewClient(serverMode, env.Env.ScrapelessStorageUrl)
	})
	return defaultStorage
}

func (s *Storage) Close() error {
	return nil
}
