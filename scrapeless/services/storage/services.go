package storage

import (
	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage"
)

type Storage struct {
	*Dataset
	*KV
	*Object
	*Queue
}

var (
	defaultStorage Storage
)

func init() {
	// todo Judge whether it is online environment IS_ONLINE according to environment variables
	defaultStorage = Storage{
		Dataset: &Dataset{},
		KV:      &KV{},
		Object:  &Object{},
		Queue:   &Queue{},
	}
}

func NewStorage(serverMode string) Storage {
	storage.NewClient(serverMode, env.Env.ScrapelessBaseApiUrl)
	return defaultStorage
}

func (s *Storage) Close() error {
	return nil
}
