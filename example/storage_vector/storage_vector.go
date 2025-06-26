package main

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/scrapeless"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/storage"
)

func main() {
	client := scrapeless.New(scrapeless.WithStorage())
	ctx := context.Background()
	collection, err := client.Storage.Vector.CreateCollections(ctx, &storage.CreateCollectionRequest{
		Name:        "my_collection",
		Description: "my collection description",
		Dimension:   1536, // Dimension of storage vector
	})

	if err != nil {
		log.Error("failed to create collection")
		return
	}
	log.Info(collection)

	createDocsResp, err := client.Storage.Vector.CreateDocs(ctx, collection.Coll.Id, []*storage.BaseDoc{
		{
			Vector: []float64{}, // The length of vector must meet the dimension of collection.
			Content: `In recent years, artificial intelligence has rapidly transformed various industries, from healthcare to finance. 
The ability of machines to learn from data and make intelligent decisions is no longer a futuristic dream, but a present-day reality. 
As we continue to integrate AI into our daily lives, questions about ethics, transparency, and human-AI collaboration become increasingly important`,
		},
	})
	if err != nil {
		log.Error("failed to push to queue")
		return
	}
	log.Info(createDocsResp)

	docs, err := client.Storage.Vector.QueryDocs(ctx, collection.Coll.Id, &storage.QueryVectorParam{
		Vector:         []float64{}, // The vector used to match the query.
		Topk:           1,           // How many results to return
		IncludeVector:  true,        // Whether to return vector
		IncludeContent: true,        // Whether to return content
	})
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Infof("query result: %+v", docs)

}
