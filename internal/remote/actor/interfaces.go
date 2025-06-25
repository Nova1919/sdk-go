package actor

import (
	"context"
	actor_http "github.com/smash-hq/sdk-go/internal/remote/actor/http"
	"github.com/smash-hq/sdk-go/internal/remote/actor/models"
	"github.com/smash-hq/sdk-go/scrapeless/log"
)

type Actor interface {
	Run(ctx context.Context, req *models.IRunActorData) (string, error)
	GetRunInfo(ctx context.Context, runId string) (*models.RunInfo, error)
	AbortRun(ctx context.Context, actorId, runId string) (bool, error)
	Build(ctx context.Context, actorId string, version string) (string, error)
	GetBuildStatus(ctx context.Context, actorId string, buildId string) (*models.BuildInfo, error)
	AbortBuild(ctx context.Context, actorId string, buildId string) (bool, error)
	GetRunList(ctx context.Context, paginationParams *models.IPaginationParams) ([]models.Payload, error)
}

var ClientInterface Actor

func NewClient(serverMode, baseUrl string) {
	switch serverMode {
	case "grpc":
		log.Info("grpc...")
	case "dev":
		log.Info("dev...")
	default:
		actor_http.Init(baseUrl)
		ClientInterface = actor_http.Default()
	}
}
