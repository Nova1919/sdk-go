package actor

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/internal/remote/actor"
)

type ActorService interface {
	Run(ctx context.Context, req actor.IRunActorData) (string, error)
	GetRunInfo(ctx context.Context, runId string) (*actor.RunInfo, error)
	AbortRun(ctx context.Context, actorId, runId string) (bool, error)
	Build(ctx context.Context, actorId string, version string) (string, error)
	GetBuildStatus(ctx context.Context, actorId string, buildId string) (*actor.BuildInfo, error)
	AbortBuild(ctx context.Context, actorId string, buildId string) (bool, error)
	GetRunList(ctx context.Context, paginationParams actor.IPaginationParams) ([]actor.Payload, error)
	Close() error
}
