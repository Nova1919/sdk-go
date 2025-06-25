package captcha

import (
	"context"
	captcha_http "github.com/smash-hq/sdk-go/internal/remote/captcha/http"
	"github.com/smash-hq/sdk-go/internal/remote/captcha/models"
	"github.com/smash-hq/sdk-go/scrapeless/log"
)

type Captcha interface {
	CaptchaSolverCreateTask(ctx context.Context, req *models.CreateTaskRequest) (string, error)
	CaptchaSolverGetTaskResult(ctx context.Context, req *models.GetTaskResultRequest) (map[string]any, error)
	CaptchaSolverSolverTask(ctx context.Context, req *models.CreateTaskRequest) (map[string]any, error)
}

var ClientInterface Captcha

func NewClient(serverMode, baseUrl string) {
	switch serverMode {
	case "grpc":
		log.Info("grpc...")
	case "dev":
		log.Info("dev...")
	default:
		captcha_http.Init(baseUrl)
		ClientInterface = captcha_http.Default()
	}
}
