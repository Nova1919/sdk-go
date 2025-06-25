package main

import (
	"context"
	"github.com/smash-hq/sdk-go/scrapeless"
	"github.com/smash-hq/sdk-go/scrapeless/log"
)

func main() {
	// new client with storage.
	client := scrapeless.New(scrapeless.WithStorage())
	defer client.Close()

	// Set value use default namespace
	ok, err := client.Storage.Kv.SetValue(context.Background(), "namespaceId", "key", "nice boy", 20)
	if err != nil {
		log.Error(err.Error())
	}
	log.Infof("ok:%v", ok)

	// Get value use default namespace
	value, err := client.Storage.Kv.GetValue(context.Background(), "namespaceId", "key")
	if err != nil {
		log.Error(err.Error())
	}
	log.Infof("value:%v", value)
}
