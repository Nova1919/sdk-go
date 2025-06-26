package universal

import (
	"context"
	"errors"
	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/internal/remote/universal"
	sh "github.com/scrapeless-ai/sdk-go/internal/remote/universal/http"
	"github.com/scrapeless-ai/sdk-go/internal/remote/universal/models"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"github.com/tidwall/gjson"
	"strings"
	"time"
)

type Universal struct{}

func New(serverMode string) *Universal {
	log.Info("Internal Universal init")
	universal.NewClient(serverMode, env.Env.ScrapelessBaseApiUrl)
	return &Universal{}
}

func (us *Universal) CreateTask(ctx context.Context, req UniversalTaskRequest) ([]byte, error) {
	if req.ProxyCountry == "" {
		req.ProxyCountry = env.Env.ProxyCountry
	}
	if req.Actor == "" {
		return nil, errors.New("actor do not be empty")
	}
	response, err := sh.Default().CreateTask(ctx, &models.UniversalTaskRequest{
		Actor: string(req.Actor),
		Input: req.Input,
		Proxy: models.TaskProxy{Country: strings.ToUpper(req.ProxyCountry)},
	})
	if err != nil {
		log.Errorf("scraping create err:%v", err)
		return nil, err
	}
	return response, nil
}

func (us *Universal) Close() error {
	return sh.Default().Close()
}

func (us *Universal) GetTaskResult(ctx context.Context, taskId string) ([]byte, error) {
	result, err := sh.Default().GetTaskResult(ctx, taskId)
	if err != nil {
		log.Errorf("get task result err:%v", err)
		return nil, err
	}
	return result, nil
}

func (us *Universal) Scrape(ctx context.Context, req UniversalTaskRequest) ([]byte, error) {
	task, err := us.CreateTask(ctx, req)
	if err != nil {
		return nil, err
	}
	taskId := gjson.Parse(string(task)).Get("taskId").String()
	if taskId != "" {
		for {
			result, err := us.GetTaskResult(ctx, taskId)
			if err == nil {
				return result, nil
			}
			time.Sleep(time.Millisecond * 200)
		}
	}
	return task, nil
}
