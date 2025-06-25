package browser

import (
	"context"
	browser_http "github.com/scrapeless-ai/sdk-go/internal/remote/browser/http"
	"github.com/scrapeless-ai/sdk-go/internal/remote/browser/models"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

type Browser interface {
	ScrapingBrowserCreate(ctx context.Context, req *models.CreateBrowserRequest) (*models.CreateBrowserResponse, error)
}

var ClientInterface Browser

func NewClient(serverMode, baseUrl string) {
	switch serverMode {
	case "grpc":
		log.Info("grpc...")
	case "dev":
		log.Info("dev...")
	default:
		browser_http.Init(baseUrl)
		ClientInterface = browser_http.Default()
	}
}
