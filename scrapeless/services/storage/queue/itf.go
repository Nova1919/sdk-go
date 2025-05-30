package queue

import (
	"context"
)

type Queue interface {
	List(ctx context.Context, page int64, pageSize int64, desc bool) (*ListQueuesResponse, error)
	Create(ctx context.Context, req *CreateQueueReq) (queueId string, queueName string, err error)
	Get(ctx context.Context, name string) (*Item, error)
	Update(ctx context.Context, name string, description string) error
	Delete(ctx context.Context) error
	pushWithId(ctx context.Context, req PushQueue) (string, error)
	Push(ctx context.Context, req PushQueue) (string, error)
	pullWithId(ctx context.Context, size int32) (GetMsgResponse, error)
	Pull(ctx context.Context, size int32) (GetMsgResponse, error)
	ackWithId(ctx context.Context, msgId string) error
	Ack(ctx context.Context, msgId string) error
	Close() error
}
