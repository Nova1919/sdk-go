package router

import (
	"github.com/scrapeless-ai/sdk-go/env"
	rh "github.com/scrapeless-ai/sdk-go/internal/remote/router/http"
	"io"

	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

type Internal struct{}

func New() Router {
	log.Info("Internal Router init")
	if rh.Default() == nil {
		rh.Init(env.Env.ScrapelessActorUrl)
	}
	return Internal{}
}

func (r Internal) Request(keyword string, method string, path string, body io.Reader, headers map[string]string) (data []byte, err error) {
	return rh.Default().Request(keyword, method, path, body, headers)
}

func (r Internal) Close() error {
	return nil
}
