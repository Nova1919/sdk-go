package storage

import (
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage/storage_http"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

type Client interface {
	Dataset
	KV
	Queue
	Object
}

var ClientInterface Client

func NewClient(serverMode, baseUrl string) {
	if serverMode == "http" {
		storage_http.Init(baseUrl)
		ClientInterface = storage_http.Default()
	} else {
		log.Info("grpc...")
	}
}
