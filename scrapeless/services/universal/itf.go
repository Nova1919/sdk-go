package universal

import "context"

type Universal interface {
	CreateTask(ctx context.Context, req UniversalTaskRequest) ([]byte, error)
	GetTaskResult(ctx context.Context, taskId string) ([]byte, error)
	Scrape(ctx context.Context, req UniversalTaskRequest) ([]byte, error)
	Close() error
}
