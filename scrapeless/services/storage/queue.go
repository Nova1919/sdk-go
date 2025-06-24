package storage

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/internal/code"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage/models"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"time"
)

type Queue struct{}

// List retrieves a list of queues with pagination and sorting options.
// Parameters:
//
//	ctx: The context for the request.
//	page: int64 - The page number (minimum 1, defaults to 1 if invalid).
//	pageSize: int64 - Number of items per page (minimum 10, defaults to 10 if invalid).
//	desc: bool - Whether to sort results in descending order.
func (s *Queue) List(ctx context.Context, page int64, pageSize int64, desc bool) (*ListQueuesResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 10 {
		pageSize = 10
	}
	queues, err := storage.ClientInterface.GetQueues(ctx, &models.GetQueuesRequest{
		Page:     page,
		PageSize: pageSize,
		Desc:     desc,
	})
	if err != nil {
		log.Errorf("failed to list queues: %v", code.Format(err))
		return nil, code.Format(err)
	}
	var items []Item
	for _, item := range queues.Items {
		items = append(items, Item{
			Id:          item.Id,
			Name:        item.Name,
			UserId:      item.UserId,
			TeamId:      item.TeamId,
			ActorId:     item.ActorId,
			RunId:       item.RunId,
			Description: item.Description,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}
	return &ListQueuesResponse{
		Items:  items,
		Total:  queues.Total,
		Limit:  queues.Limit,
		Offset: queues.Offset,
	}, nil
}

// Create creates a new HTTP queue with the provided request parameters.
// Parameters:
//
//	ctx: The context for the request.
//	req: The request object containing queue configuration details.
func (s *Queue) Create(ctx context.Context, req *CreateQueueReq) (queueId string, queueName string, err error) {
	name := req.Name + "-" + env.GetActorEnv().RunId
	queue, err := storage.ClientInterface.CreateQueue(ctx, &models.CreateQueueRequest{
		ActorId:     env.GetActorEnv().ActorId,
		RunId:       env.GetActorEnv().RunId,
		Name:        name,
		Description: req.Description,
	})
	if err != nil {
		log.Errorf("failed to create queue: %v", code.Format(err))
		return "", "", code.Format(err)
	}

	return queue.Id, name, nil
}

// Get retrieves a queue item by name.
// Parameters:
//
//	ctx: The context for the request.
//	name: The name of the queue to retrieve.
func (s *Queue) Get(ctx context.Context, queueId string, name string) (*Item, error) {
	name = name + "-" + env.GetActorEnv().RunId
	queue, err := storage.ClientInterface.GetQueue(ctx, &models.GetQueueRequest{
		Id:   queueId,
		Name: name,
	})
	if err != nil {
		log.Errorf("failed to get queue: %v", code.Format(err))
		return nil, code.Format(err)
	}
	return &Item{
		Id:          queue.Id,
		Name:        queue.Name,
		TeamId:      queue.TeamId,
		ActorId:     queue.ActorId,
		RunId:       queue.RunId,
		Description: queue.Description,
		CreatedAt:   queue.CreatedAt,
		UpdatedAt:   queue.UpdatedAt,
	}, nil
}

// Update updates the queue information with the provided name and description.
// Parameters:
//
//	ctx: The context for the request.
//	name: The new name of the queue.
//	description: The new description of the queue.
func (s *Queue) Update(ctx context.Context, queueId string, name string, description string) error {
	name = name + "-" + env.GetActorEnv().RunId
	err := storage.ClientInterface.UpdateQueue(ctx, &models.UpdateQueueRequest{
		QueueId:     queueId,
		Name:        name,
		Description: description,
	})
	return err
}

// Delete deletes the queue using the storage HTTP service.
// Parameters:
//
//	ctx: The context for the request.
func (s *Queue) Delete(ctx context.Context, queueId string) error {
	err := storage.ClientInterface.DelQueue(ctx, &models.DelQueueRequest{QueueId: queueId})
	if err != nil {
		log.Errorf("failed to delete queue: %v", code.Format(err))
		return code.Format(err)
	}
	return nil
}

// Push adds a request to the HTTP queue and returns the task ID.  timeout-->[60,300]   deadline--> [300,86400]
//
// Parameters:
//
//	ctx context.Context: The context for the request, used for cancellation and timeouts.
//	req PushQueue: The request to be pushed into the queue.
func (s *Queue) Push(ctx context.Context, queueId string, req PushQueue) (string, error) {
	// [60,300]
	if req.Timeout < 60 {
		req.Timeout = 60
	}
	if req.Timeout > 300 {
		req.Timeout = 300
	}

	// [300,86400]
	if req.Deadline < 300 {
		req.Deadline = 400
	}
	if req.Deadline > 86400 {
		req.Deadline = 86400
	}

	unix := time.Now().UTC().Add(time.Duration(req.Deadline) * time.Second).Unix()
	queue, err := storage.ClientInterface.CreateMsg(ctx, &models.CreateMsgRequest{
		QueueId:  queueId,
		Name:     req.Name,
		PayLoad:  string(req.Payload),
		Retry:    req.Retry,
		Timeout:  req.Timeout,
		Deadline: unix,
	})
	if err != nil {
		log.Errorf("failed to push to queue: %v", code.Format(err))
		return "", code.Format(err)
	}
	return queue.MsgId, nil
}

// Pull retrieves messages from the HTTP queue.
// Parameters:
//
//	ctx: The context used to control the request lifecycle (e.g., cancellation, deadlines).
//	size: The maximum number of messages to retrieve in this operation.
func (s *Queue) Pull(ctx context.Context, queueId string, size int32) (GetMsgResponse, error) {
	if size < 1 {
		size = 1
	}
	if size > 100 {
		size = 100
	}
	msgs, err := storage.ClientInterface.GetMsg(ctx, &models.GetMsgRequest{
		QueueId: queueId,
		Limit:   size,
	})
	if err != nil {
		log.Errorf("failed to pull from queue: %v", code.Format(err))
		return nil, code.Format(err)
	}
	if msgs == nil {
		return nil, nil
	}
	var items []*Msg
	for _, msg := range *msgs {
		items = append(items, &Msg{
			ID:        msg.ID,
			QueueID:   msg.QueueID,
			Name:      msg.Name,
			Payload:   msg.Payload,
			Timeout:   msg.Timeout,
			Deadline:  msg.Deadline,
			Retry:     msg.Retry,
			Retried:   msg.Retried,
			SuccessAt: msg.SuccessAt,
			FailedAt:  msg.FailedAt,
			Desc:      msg.Desc,
		})
	}
	return items, nil
}

// Ack confirms that a message has been processed successfully.
//
// Parameters:
//
//	ctx: The context used for request cancellation or timeout.
//	msgId: The unique identifier of the message to acknowledge.
func (s *Queue) Ack(ctx context.Context, queueId string, msgId string) error {
	err := storage.ClientInterface.AckMsg(ctx, &models.AckMsgRequest{
		QueueId: queueId,
		MsgId:   msgId,
	})
	if err != nil {
		log.Errorf("failed to ack msg: %v", code.Format(err))
		return code.Format(err)
	}
	return nil
}

func (s *Queue) Close() error {
	return nil
}
