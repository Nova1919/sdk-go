package browser

import (
	"context"
)

type Browser interface {
	Create(ctx context.Context, req Actor) (*CreateResp, error)
	CreateOnce(ctx context.Context, req ActorOnce) (*CreateResp, error)
	Close() error
}
