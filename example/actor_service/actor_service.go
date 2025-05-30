package main

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/internal/remote/actor"
	"github.com/scrapeless-ai/sdk-go/scrapeless"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

func main() {
	client := scrapeless.New(scrapeless.WithActor())
	defer client.Close()

	runId, err := client.Actor.Run(context.Background(), actor.IRunActorData{
		ActorId: "554bbd68-c787-4900-b8b2-1086369c96e1",
		Input: map[string]string{
			"name": "test",
			"url":  "https://www.google.com",
		},
		RunOptions: actor.RunOptions{
			Version: "v0.0.3",
		},
	})
	if err != nil {
		panic(err)
	}
	runInfo, err := client.Actor.GetRunInfo(context.Background(), runId)
	if err != nil {
		panic(err)
	}
	log.Infof("runInfo:%+v", runInfo)
}
