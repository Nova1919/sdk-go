package main

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/scrapeless"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/storage/queue"
)

func main() {
	client := scrapeless.New(scrapeless.WithStorage())

	// push a message to queue
	msgId, err := client.Storage.Queue.Push(context.Background(), "queueId", queue.PushQueue{
		Name:    "test-cy",
		Payload: []byte("aaaa"),
	})
	if err != nil {
		log.Error("failed to push to queue")
		return
	}
	log.Info(msgId)

	// pull a message from queue
	pullResp, err := client.Storage.Queue.Pull(context.Background(), "queueId", 100)
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Infof("%v", pullResp)
	for _, v := range pullResp {
		// ack message
		err = client.Storage.Queue.Ack(context.Background(), "queueId", v.QueueID)
		if err != nil {
			log.Error(err.Error())
		}
	}

}
