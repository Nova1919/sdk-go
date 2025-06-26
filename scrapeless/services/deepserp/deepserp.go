package deepserp

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/internal/code"
	"github.com/scrapeless-ai/sdk-go/internal/remote/deepserp"
	dh "github.com/scrapeless-ai/sdk-go/internal/remote/deepserp/http"
	"github.com/scrapeless-ai/sdk-go/internal/remote/deepserp/models"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"github.com/tidwall/gjson"
	"strings"
	"time"
)

type DeepSerp struct{}

func NewDeepSerp(serverMode string) *DeepSerp {
	log.Info("Internal DeepSerp init")
	deepserp.NewClient(serverMode, env.Env.ScrapelessBaseApiUrl)
	return &DeepSerp{}
}

// CreateTask creates a new deepSerp task with the given context and request parameters.
func (s *DeepSerp) CreateTask(ctx context.Context, req DeepserpTaskRequest) ([]byte, error) {
	if req.ProxyCountry == "" {
		req.ProxyCountry = env.Env.ProxyCountry
	}
	response, err := deepserp.ClientInterface.CreateTask(ctx, &models.DeepserpTaskRequest{
		Actor: string(req.Actor),
		Input: req.Input,
		Proxy: models.TaskProxy{Country: strings.ToUpper(req.ProxyCountry)},
	})
	if err != nil {
		log.Errorf("deepserp create err:%v", err)
		return nil, code.Format(err)
	}
	return response, nil
}

func (s *DeepSerp) Close() error {
	return dh.Default().Close()
}

// GetTaskResult retrieves the result of a deepSerp task by its ID.
func (s *DeepSerp) GetTaskResult(ctx context.Context, taskId string) ([]byte, error) {
	result, err := deepserp.ClientInterface.GetTaskResult(ctx, taskId)
	if err != nil {
		log.Errorf("get task result err:%v", err)
		return nil, code.Format(err)
	}
	return result, nil
}

func (s *DeepSerp) Scrape(ctx context.Context, req DeepserpTaskRequest) ([]byte, error) {
	task, err := s.CreateTask(ctx, req)
	if err != nil {
		return nil, err
	}
	taskId := gjson.Parse(string(task)).Get("taskId").String()
	if taskId != "" {
		for {
			result, err := s.GetTaskResult(ctx, taskId)
			if err == nil {
				return result, nil
			}
			time.Sleep(time.Millisecond * 200)
		}
	}
	return task, nil
}
