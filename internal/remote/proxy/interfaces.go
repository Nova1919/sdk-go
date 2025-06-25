package proxy

import (
	"context"
	proxy_http "github.com/scrapeless-ai/sdk-go/internal/remote/proxy/http"
	"github.com/scrapeless-ai/sdk-go/internal/remote/proxy/models"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

type Proxy interface {
	ProxyGetProxy(ctx context.Context, req *models.GetProxyRequest) (string, error)
}

var ClientInterface Proxy

func NewClient(serverMode string) {
	switch serverMode {
	case "grpc":
		log.Info("grpc...")
	case "dev":
		log.Info("dev...")
	default:
		proxy_http.Init()
		ClientInterface = proxy_http.Default()
	}
}
