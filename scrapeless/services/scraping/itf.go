package scraping

import "context"

type Scraping interface {
	CreateTask(ctx context.Context, req ScrapingTaskRequest) ([]byte, error)
	GetTaskResult(ctx context.Context, taskId string) ([]byte, error)
	Scrape(ctx context.Context, req ScrapingTaskRequest) ([]byte, error)
	Close() error
}
