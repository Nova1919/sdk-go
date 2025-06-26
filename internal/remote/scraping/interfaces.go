package scraping

import (
	"context"
	scraping_http "github.com/scrapeless-ai/sdk-go/internal/remote/scraping/http"
	"github.com/scrapeless-ai/sdk-go/internal/remote/scraping/models"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

type Scraping interface {
	Scrape(ctx context.Context, req *models.ScrapingRequest) ([]byte, error)
	CreateTask(ctx context.Context, req *models.ScrapingTaskRequest) ([]byte, error)
	GetTaskResult(ctx context.Context, taskIKd string) ([]byte, error)
}

var ClientInterface Scraping

func NewClient(serverMode, baseUrl string) {
	switch serverMode {
	case "grpc":
		log.Info("grpc...")
	case "dev":
		log.Info("dev...")
	default:
		scraping_http.Init(baseUrl)
		ClientInterface = scraping_http.Default()
	}
}
