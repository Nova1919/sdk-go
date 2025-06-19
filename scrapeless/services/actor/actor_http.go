package actor

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/internal/code"
	"github.com/scrapeless-ai/sdk-go/internal/remote/actor"
	actor_http "github.com/scrapeless-ai/sdk-go/internal/remote/actor/http"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

func NewActorHttp() ActorService {
	log.Info("Actor http init")
	if actor_http.Default() == nil {
		actor_http.Init(env.Env.ScrapelessActorUrl)
	}
	return &ActorHttp{}
}

type ActorHttp struct {
}

// Run starts an actor run with the provided context and request data.
// Returns the run ID or an error.
func (ah *ActorHttp) Run(ctx context.Context, req actor.IRunActorData) (string, error) {
	runId, err := actor_http.Default().Run(ctx, req)
	return runId, code.Format(err)
}

// GetRunInfo retrieves information about a specific actor run by run ID.
// Returns a pointer to RunInfo or an error.
func (ah *ActorHttp) GetRunInfo(ctx context.Context, runId string) (*actor.RunInfo, error) {
	runInfo, err := actor_http.Default().GetRunInfo(ctx, runId)
	return runInfo, code.Format(err)
}

// AbortRun aborts a running actor by actor ID and run ID.
// Returns true if successful and an error otherwise.
func (ah *ActorHttp) AbortRun(ctx context.Context, actorId, runId string) (bool, error) {
	success, err := actor_http.Default().AbortRun(ctx, actorId, runId)
	return success, code.Format(err)
}

// Build triggers a build process for the specified actor and version.
// Returns the build ID or an error.
func (ah *ActorHttp) Build(ctx context.Context, actorId string, version string) (string, error) {
	buildId, err := actor_http.Default().Build(ctx, actorId, version)
	return buildId, code.Format(err)
}

// GetBuildStatus retrieves the status of a build by actor ID and build ID.
// Returns a pointer to BuildInfo or an error.
func (ah *ActorHttp) GetBuildStatus(ctx context.Context, actorId string, buildId string) (*actor.BuildInfo, error) {
	success, err := actor_http.Default().GetBuildStatus(ctx, actorId, buildId)
	return success, code.Format(err)
}

// AbortBuild aborts an ongoing build process by actor ID and build ID.
// Returns true if successful and an error otherwise.
func (ah *ActorHttp) AbortBuild(ctx context.Context, actorId string, buildId string) (bool, error) {
	success, err := actor_http.Default().AbortBuild(ctx, actorId, buildId)
	return success, code.Format(err)
}

// GetRunList retrieves a list of actor runs with pagination.
// Returns a slice of Payload containing run data or an error.
func (ah *ActorHttp) GetRunList(ctx context.Context, paginationParams actor.IPaginationParams) ([]actor.Payload, error) {
	runList, err := actor_http.Default().GetRunList(ctx, paginationParams)
	return runList, code.Format(err)
}

func (ah *ActorHttp) Close() error {
	return actor_http.Default().Close()
}
