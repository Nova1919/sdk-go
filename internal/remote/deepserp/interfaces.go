package deepserp

import (
	"context"
	deepserp_http "github.com/smash-hq/sdk-go/internal/remote/deepserp/http"
	"github.com/smash-hq/sdk-go/internal/remote/deepserp/models"
	"github.com/smash-hq/sdk-go/scrapeless/log"
)

type DeepSerp interface {
	CreateTask(ctx context.Context, req *models.DeepserpTaskRequest) ([]byte, error)
	GetTaskResult(ctx context.Context, taskIKd string) ([]byte, error)
}

var ClientInterface DeepSerp

func NewClient(serverMode, baseUrl string) {
	switch serverMode {
	case "grpc":
		log.Info("grpc...")
	case "dev":
		log.Info("dev...")
	default:
		deepserp_http.Init(baseUrl)
		ClientInterface = deepserp_http.Default()
	}
}
