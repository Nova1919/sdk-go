package universal

import (
	"context"
	"errors"
	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/internal/remote/universal"
	sh "github.com/scrapeless-ai/sdk-go/internal/remote/universal/http"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"github.com/tidwall/gjson"
	"strings"
	"time"
)

type UniversalHttp struct{}

func New() Universal {
	log.Info("Internal Router init")
	if sh.Default() == nil {
		sh.Init(env.Env.ScrapelessBaseApiUrl)
	}
	return UniversalHttp{}
}

func (us UniversalHttp) CreateTask(ctx context.Context, req UniversalTaskRequest) ([]byte, error) {
	if req.ProxyCountry == "" {
		req.ProxyCountry = env.Env.ProxyCountry
	}
	if req.Actor == "" {
		return nil, errors.New("actor do not be empty")
	}
	response, err := sh.Default().CreateTask(ctx, universal.UniversalTaskRequest{
		Actor: string(req.Actor),
		Input: req.Input,
		Proxy: universal.TaskProxy{Country: strings.ToUpper(req.ProxyCountry)},
	})
	if err != nil {
		log.Errorf("scraping create err:%v", err)
		return nil, err
	}
	return response, nil
}

func (us UniversalHttp) Close() error {
	return sh.Default().Close()
}

func (us UniversalHttp) GetTaskResult(ctx context.Context, taskId string) ([]byte, error) {
	result, err := sh.Default().GetTaskResult(ctx, taskId)
	if err != nil {
		log.Errorf("get task result err:%v", err)
		return nil, err
	}
	return result, nil
}

func (us UniversalHttp) Scrape(ctx context.Context, req UniversalTaskRequest) ([]byte, error) {
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
