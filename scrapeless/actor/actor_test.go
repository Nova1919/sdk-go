package actor

import (
	"context"
	proxy2 "github.com/smash-hq/sdk-go/scrapeless/services/proxies"
	"testing"
)

func TestNew(t *testing.T) {
	actor := New()
	p, _ := actor.Proxy.Proxy(context.Background(), proxy2.ProxyActor{
		Country:         "",
		SessionDuration: 0,
		SessionId:       "",
		Gateway:         "",
	})
	t.Log(p)
}
