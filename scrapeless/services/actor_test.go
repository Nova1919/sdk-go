package services

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/actor"
	"testing"
)

func TestActor(t *testing.T) {
	ah := actor.NewActor("http")
	run, err := ah.Run(context.Background(), actor.IRunActorData{
		ActorId: "554bbd68-c787-4900-b8b2-1086369c96e1",
		Input: map[string]string{
			"name": "test",
			"url":  "https://www.google.com",
		},
		RunOptions: actor.RunOptions{
			Version: "v0.0.3",
		},
	})
	if err != nil {
		t.Errorf("run err:%v", err)
		return
	}
	t.Logf("run:%v", run)
}

func TestGetRunInfo(t *testing.T) {
	ah := actor.NewActor("http")
	runInfo, err := ah.GetRunInfo(context.Background(), "2efd2e25-c6aa-4934-b097-8b3b78dda451")
	if err != nil {
		t.Errorf("run err:%v", err)
		return
	}
	t.Logf("run:%v", runInfo)
}

func TestAbortRun(t *testing.T) {
	ah := actor.NewActor("http")
	ok, err := ah.AbortRun(context.Background(), "554bbd68-c787-4900-b8b2-1086369c96e1", "2efd2e25-c6aa-4934-b097-8b3b78dda451")
	if err != nil {
		t.Errorf("run err:%v", err)
		return
	}
	t.Logf("run:%v", ok)
}
func TestGetRunList(t *testing.T) {
	ah := actor.NewActor("http")
	runList, err := ah.GetRunList(context.Background(), &actor.IPaginationParams{
		Page:     1,
		PageSize: 1,
	})
	if err != nil {
		t.Errorf("run err:%v", err)
		return
	}
	t.Logf("run:%+v", runList[0])
}
func TestBuild(t *testing.T) {
	ah := actor.NewActor("http")
	buildId, err := ah.Build(context.Background(), "0420153f-c002-4417-94ac-b9135ea22ae4", "main")
	if err != nil {
		t.Errorf("run err:%v", err)
		return
	}
	t.Logf("run:%+v", buildId)
}

func TestGetBuildStatus(t *testing.T) {
	ah := actor.NewActor("http")

	ok, err := ah.GetBuildStatus(context.Background(), "0420153f-c002-4417-94ac-b9135ea22ae4", "955ec451-4e70-4283-a106-31669f6f08bb")
	if err != nil {
		t.Errorf("run err:%v", err)
		return
	}
	t.Logf("run:%+v", ok)
}
