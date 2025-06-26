package universal

import (
	"context"
	universal_http "github.com/scrapeless-ai/sdk-go/internal/remote/universal/http"
	"github.com/scrapeless-ai/sdk-go/internal/remote/universal/models"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

type Browser interface {
	CreateTask(ctx context.Context, req *models.UniversalTaskRequest) ([]byte, error)
	GetTaskResult(ctx context.Context, taskIKd string) ([]byte, error)
}

var ClientInterface Browser

func NewClient(serverMode, baseUrl string) {
	switch serverMode {
	case "grpc":
		log.Info("grpc...")
	case "dev":
		log.Info("dev...")
	default:
		universal_http.Init(baseUrl)
		ClientInterface = universal_http.Default()
	}
}
