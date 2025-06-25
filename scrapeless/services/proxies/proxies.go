package proxies

import (
	"context"
	"github.com/smash-hq/sdk-go/env"
	"github.com/smash-hq/sdk-go/internal/code"
	rp "github.com/smash-hq/sdk-go/internal/remote/proxy"
	"github.com/smash-hq/sdk-go/internal/remote/proxy/http"
	proxy2 "github.com/smash-hq/sdk-go/internal/remote/proxy/models"
	"github.com/smash-hq/sdk-go/scrapeless/log"
)

type Proxy struct {
}

func NewProxy(serverMode string) *Proxy {
	log.Infof("proxies init")
	rp.NewClient(serverMode)
	return &Proxy{}
}

// Proxy retrieves proxies information.
//
// Parameters:
//
//	ctx: context.Context - Context for the request.
//	proxies: ProxyActor - Struct containing proxies request parameters like country, session duration, etc.
func (ph *Proxy) Proxy(ctx context.Context, proxy ProxyActor) (string, error) {
	proxyUrl, err := rp.ClientInterface.ProxyGetProxy(ctx, &proxy2.GetProxyRequest{
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

func (ph *Proxy) Close() error {
	return http.Default().Close()
}
