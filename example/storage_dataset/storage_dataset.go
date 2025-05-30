package main

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/scrapeless"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

func main() {
	client := scrapeless.New(scrapeless.WithStorage())
	defer client.Close()

	success, err := client.Storage.GetDataset().AddItems(context.Background(), []map[string]any{
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
		items, err := client.Storage.GetDataset().GetItems(context.Background(), 1, 10, false)
		if err != nil {
			log.Error(err.Error())
			return
		}
		log.Infof("%v", items)
	}

}
