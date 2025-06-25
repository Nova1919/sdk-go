package router

import (
	router_http "github.com/smash-hq/sdk-go/internal/remote/router/http"
	"github.com/smash-hq/sdk-go/scrapeless/log"
	"io"
)

type Router interface {
	Request(keyword string, method string, path string, body io.Reader, headers map[string]string) (data []byte, err error)
}

var ClientInterface Router

func NewClient(serverMode string) {
	switch serverMode {
	case "grpc":
		log.Info("grpc...")
	case "dev":
		log.Info("dev...")
	default:
		router_http.Init()
		ClientInterface = router_http.Default()
	}
}
