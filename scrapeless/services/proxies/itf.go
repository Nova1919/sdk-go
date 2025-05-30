package proxies

import "context"

type Proxy interface {
	Proxy(ctx context.Context, proxy ProxyActor) (string, error)
	Close() error
}
