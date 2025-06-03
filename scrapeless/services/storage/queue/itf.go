package queue

import (
	"context"
)

type Queue interface {
	List(ctx context.Context, page int64, pageSize int64, desc bool) (*ListQueuesResponse, error)
	Create(ctx context.Context, req *CreateQueueReq) (queueId string, queueName string, err error)
	Get(ctx context.Context, queueId string, name string) (*Item, error)
	Update(ctx context.Context, queueId string, name string, description string) error
	Delete(ctx context.Context, queueId string) error
	Push(ctx context.Context, queueId string, req PushQueue) (string, error)
	Pull(ctx context.Context, queueId string, size int32) (GetMsgResponse, error)
	Ack(ctx context.Context, queueId string, msgId string) error
	Close() error
}
