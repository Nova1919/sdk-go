package storage

import (
	"github.com/smash-hq/sdk-go/env"
	"github.com/smash-hq/sdk-go/internal/remote/storage"
	"sync"
)

type Storage struct {
	*Dataset
	*KV
	*Object
	*Queue
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
