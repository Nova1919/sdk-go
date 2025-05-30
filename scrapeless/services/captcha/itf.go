package captcha

import "context"

type Captcha interface {
	Solver(ctx context.Context, req *CaptchaSolverReq) (*CaptchaSolverResp, error)
	Create(ctx context.Context, req *CaptchaSolverReq) (string, error)
	ResultGet(ctx context.Context, req *CaptchaSolverReq) (*CaptchaSolverResp, error)
	Close() error
}
