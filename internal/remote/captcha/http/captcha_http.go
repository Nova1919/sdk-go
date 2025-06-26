package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/scrapeless-ai/sdk-go/internal/remote/captcha/models"
	request2 "github.com/scrapeless-ai/sdk-go/internal/remote/request"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"github.com/tidwall/gjson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

func (c *Client) CaptchaSolverCreateTask(ctx context.Context, req *models.CreateTaskRequest) (string, error) {
	reqBodyM := map[string]any{
		"service": req.Actor,
		"input":   req.Input,
		"proxies": req.Proxy,
	}
	reqBody, err := json.Marshal(reqBodyM)
	if err != nil {
		return "", err
	}
	fmt.Println(fmt.Sprintf("%s/api/v1/createTask", c.BaseUrl))
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodPost,
		Url:    fmt.Sprintf("%s/api/v1/createTask", c.BaseUrl),
		Body:   string(reqBody),
		Headers: map[string]string{
			"x-api-key": req.ApiKey,
			"token":     req.ApiKey,
		},
	})
	if err != nil {
		return "", err
	}
	taskId := gjson.Parse(body).Get("taskId").String()
	if taskId == "" {
		msg := gjson.Parse(body).Get("message").String()
		return "", fmt.Errorf("create task err:%s", msg)
	}
	return taskId, nil

}

func (c *Client) CaptchaSolverGetTaskResult(ctx context.Context, req *models.GetTaskResultRequest) (map[string]any, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("%s/api/v1/getTaskResult/%s", c.BaseUrl, req.TaskId),
		Headers: map[string]string{
			"x-api-key": req.ApiKey,
			"token":     req.ApiKey,
		},
	})
	if err != nil {
		return nil, err
	}
	if ok := gjson.Parse(body).Get("success").Bool(); !ok {
		log.Error(body)
		return nil, fmt.Errorf("get task result err")
	}
	var solution map[string]any
	solutionStr := gjson.Parse(body).Get("solution").String()
	if err = json.Unmarshal([]byte(solutionStr), &solution); err != nil {
		return nil, err
	}
	return solution, nil
}

func (c *Client) CaptchaSolverSolverTask(ctx context.Context, req *models.CreateTaskRequest) (map[string]any, error) {
	task, err := c.CaptchaSolverCreateTask(ctx, req)
	if err != nil {
		return nil, err
	}
	for {
		select {
		case <-ctx.Done():
			return nil, status.Errorf(codes.DeadlineExceeded, ctx.Err().Error())
		case <-time.After(time.Second):
			result, err := c.CaptchaSolverGetTaskResult(ctx, &models.GetTaskResultRequest{TaskId: task, ApiKey: req.ApiKey})
			if err != nil {
				return nil, status.Errorf(codes.Aborted, err.Error())
			}
			return result, nil
		}
	}
}
