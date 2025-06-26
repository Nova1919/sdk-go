package router

import (
	"github.com/scrapeless-ai/sdk-go/internal/remote/router"
	"io"

	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

type Router struct{}

func New(serverMode string) *Router {
	log.Info("Internal Router init")
	router.NewClient(serverMode)
	return &Router{}
}

// Request keyword is the actor's keyword-->Now its value is runnerId
func (r *Router) Request(keyword string, method string, path string, body io.Reader, headers map[string]string) (data []byte, err error) {
	return router.ClientInterface.Request(keyword, method, path, body, headers)
}

func (r *Router) Close() error {
	return nil
}
