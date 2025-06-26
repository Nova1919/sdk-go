package main

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/scrapeless"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

func main() {
	client := scrapeless.New(scrapeless.WithStorage())
	defer client.Close()

	// Put object The supported types include JSON、html、png
	objectId, err := client.Storage.Object.Put(context.Background(), "bucketId", "object.json", []byte("data"))
	if err != nil {
		log.Error(err.Error())
		return
	}
	if objectId != "" {
		// Get object
		resp, err := client.Storage.Object.Get(context.Background(), "bucketId", objectId)
		if err != nil {
			panic(err)
		}
		log.Info(string(resp))
	}
}
