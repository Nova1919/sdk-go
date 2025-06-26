package main

import (
	"context"
	"github.com/smash-hq/sdk-go/scrapeless"
	"github.com/smash-hq/sdk-go/scrapeless/log"
)

func main() {
	client := scrapeless.New(scrapeless.WithStorage())
	defer client.Close()

	success, err := client.Storage.Dataset.AddItems(context.Background(), "datasetId", []map[string]any{
		{
			"name": "John",
			"age":  20,
		},
		{
			"name": "lucy",
			"age":  19,
		},
	})
	if err != nil {
		log.Error(err.Error())
		return
	}
	if success {
		items, err := client.Storage.Dataset.GetItems(context.Background(), "datasetId", 1, 10, false)
		if err != nil {
			log.Error(err.Error())
			return
		}
		log.Infof("%v", items)
	}

}
