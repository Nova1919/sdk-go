package browser

import (
	"context"
)

type BGrpc struct {
}

func NewBGrpc() Browser {
	return &BGrpc{}
}

func (bg *BGrpc) Create(ctx context.Context, req Actor) (*CreateResp, error) {
	return nil, nil
}

func (bg *BGrpc) CreateOnce(ctx context.Context, req ActorOnce) (*CreateResp, error) {
	return nil, nil
}
func (bg *BGrpc) Close() error {
	return nil
}
