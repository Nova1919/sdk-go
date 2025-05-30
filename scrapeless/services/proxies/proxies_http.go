package proxies

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/internal/code"
	proxy2 "github.com/scrapeless-ai/sdk-go/internal/remote/proxy"
	"github.com/scrapeless-ai/sdk-go/internal/remote/proxy/http"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

type PHttp struct {
}

func NewPHttp() Proxy {
	log.Infof("proxies http init")
	if http.Default() == nil {
		http.Init()
	}
	return &PHttp{}
}

// Proxy retrieves proxies information.
//
// Parameters:
//
//	ctx: context.Context - Context for the request.
//	proxies: ProxyActor - Struct containing proxies request parameters like country, session duration, etc.
func (ph *PHttp) Proxy(ctx context.Context, proxy ProxyActor) (string, error) {
	proxyUrl, err := http.Default().ProxyGetProxy(ctx, &proxy2.GetProxyRequest{
		ApiKey:          env.GetActorEnv().ApiKey,
		Country:         proxy.Country,
		SessionDuration: proxy.SessionDuration,
		SessionId:       proxy.SessionId,
		Gateway:         proxy.Gateway,
		TaskId:          env.GetActorEnv().RunId,
	})
	if err != nil {
		log.Errorf("get proxies err:%v", err)
		return "", code.Format(err)
	}
	return proxyUrl, nil
}

func (ph *PHttp) Close() error {
	return http.Default().Close()
}
