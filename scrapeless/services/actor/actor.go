package actor

import (
	"context"
	"github.com/smash-hq/sdk-go/env"
	"github.com/smash-hq/sdk-go/internal/code"
	"github.com/smash-hq/sdk-go/internal/remote/actor"
	actor_http "github.com/smash-hq/sdk-go/internal/remote/actor/http"
	"github.com/smash-hq/sdk-go/internal/remote/actor/models"
	"github.com/smash-hq/sdk-go/scrapeless/log"
)

func NewActor(serverMode string) *ActorService {
	log.Info("Actor init")
	actor.NewClient(serverMode, env.Env.ScrapelessActorUrl)
	return &ActorService{}
}

type ActorService struct {
}

// Run starts an actor run with the provided context and request data.
// Returns the run ID or an error.
func (ah *ActorService) Run(ctx context.Context, req IRunActorData) (string, error) {
	runId, err := actor_http.Default().Run(ctx, &models.IRunActorData{
		ActorId: req.ActorId,
		Input:   req.Input,
		RunOptions: models.RunOptions{
			CPU:     req.RunOptions.CPU,
			Memory:  req.RunOptions.Memory,
			Timeout: req.RunOptions.Timeout,
			Version: req.RunOptions.Version,
		},
	})
	return runId, code.Format(err)
}

// GetRunInfo retrieves information about a specific actor run by run ID.
// Returns a pointer to RunInfo or an error.
func (ah *ActorService) GetRunInfo(ctx context.Context, runId string) (*RunInfo, error) {
	runInfo, err := actor_http.Default().GetRunInfo(ctx, runId)
	if err != nil {
		log.Errorf("get runInfo err:%v", err)
		return nil, code.Format(err)
	}
	info := &RunInfo{
		ActorID:     runInfo.ActorID,
		ActorName:   runInfo.ActorName,
		FinishedAt:  runInfo.FinishedAt,
		Input:       runInfo.Input,
		InputSchema: runInfo.InputSchema,
		Origin:      runInfo.Origin,
		RunID:       runInfo.RunID,
		RunOptions: ResourceOptions{
			CPU:    runInfo.RunOptions.CPU,
			Memory: runInfo.RunOptions.Memory,
		},
		SchedulerID: runInfo.SchedulerID,
		StartedAt:   runInfo.StartedAt,
		Status:      runInfo.Status,
		Storage: StorageInfo{
			BucketID:      runInfo.Storage.BucketID,
			DatasetID:     runInfo.Storage.DatasetID,
			KVNamespaceID: runInfo.Storage.KVNamespaceID,
			QueueID:       runInfo.Storage.QueueID,
		},
		TeamID: runInfo.TeamID,
	}
	return info, nil
}

// AbortRun aborts a running actor by actor ID and run ID.
// Returns true if successful and an error otherwise.
func (ah *ActorService) AbortRun(ctx context.Context, actorId, runId string) (bool, error) {
	success, err := actor_http.Default().AbortRun(ctx, actorId, runId)
	return success, code.Format(err)
}

// Build triggers a build process for the specified actor and version.
// Returns the build ID or an error.
func (ah *ActorService) Build(ctx context.Context, actorId string, version string) (string, error) {
	buildId, err := actor_http.Default().Build(ctx, actorId, version)
	return buildId, code.Format(err)
}

// GetBuildStatus retrieves the status of a build by actor ID and build ID.
// Returns a pointer to BuildInfo or an error.
func (ah *ActorService) GetBuildStatus(ctx context.Context, actorId string, buildId string) (*BuildInfo, error) {
	success, err := actor_http.Default().GetBuildStatus(ctx, actorId, buildId)
	if err != nil {
		log.Errorf("get build status err:%v", err)
		return nil, code.Format(err)
	}
	buildInfo := &BuildInfo{
		ActorID:    success.ActorID,
		BuildID:    success.BuildID,
		Duration:   success.Duration,
		FinishedAt: success.FinishedAt,
		ImageSize:  success.ImageSize,
		RepoID:     success.RepoID,
		StartedAt:  success.StartedAt,
		Status:     success.Status,
		TeamID:     success.TeamID,
		Version:    success.Version,
	}
	return buildInfo, code.Format(err)
}

// AbortBuild aborts an ongoing build process by actor ID and build ID.
// Returns true if successful and an error otherwise.
func (ah *ActorService) AbortBuild(ctx context.Context, actorId string, buildId string) (bool, error) {
	success, err := actor_http.Default().AbortBuild(ctx, actorId, buildId)
	return success, code.Format(err)
}

// GetRunList retrieves a list of actor runs with pagination.
// Returns a slice of Payload containing run data or an error.
func (ah *ActorService) GetRunList(ctx context.Context, paginationParams *IPaginationParams) ([]Payload, error) {
	runList, err := actor_http.Default().GetRunList(ctx, &models.IPaginationParams{
		Page:     paginationParams.Page,
		PageSize: paginationParams.PageSize,
		Desc:     paginationParams.Desc,
	})
	if err != nil {
		log.Errorf("get run list err:%v", err)
		return nil, code.Format(err)
	}
	var runListArray []Payload
	for _, run := range runList {
		runListArray = append(runListArray, Payload{
			ActorID:     run.ActorID,
			ActorName:   run.ActorName,
			FinishedAt:  run.FinishedAt,
			Input:       run.Input,
			InputSchema: run.InputSchema,
			Origin:      run.Origin,
			RunID:       run.RunID,
			RunOptions: ResourceOptions{
				CPU:          run.RunOptions.CPU,
				Memory:       run.RunOptions.Memory,
				ServerMode:   run.RunOptions.ServerMode,
				SurvivalTime: run.RunOptions.SurvivalTime,
				Timeout:      run.RunOptions.Timeout,
				Version:      run.RunOptions.Version,
			},
			SchedulerID: run.SchedulerID,
			StartedAt:   run.StartedAt,
			Status:      run.Status,
			Storage: StorageInfo{
				BucketID:      run.Storage.BucketID,
				DatasetID:     run.Storage.DatasetID,
				KVNamespaceID: run.Storage.KVNamespaceID,
				QueueID:       run.Storage.QueueID,
			},
			TeamID: run.TeamID,
		})
	}
	return runListArray, nil
}

func (ah *ActorService) Close() error {
	return actor_http.Default().Close()
}
