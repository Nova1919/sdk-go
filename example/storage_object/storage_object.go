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
	objectId, err := client.Storage.GetObject().Put(context.Background(), "object.json", []byte("data"))
	if err != nil {
		log.Error(err.Error())
		return
	}
	if objectId != "" {
		// Get object
		resp, err := client.Storage.GetObject().Get(context.Background(), objectId)
		if err != nil {
			panic(err)
		}
		log.Info(string(resp))
	}
}
