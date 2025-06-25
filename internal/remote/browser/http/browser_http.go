package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/smash-hq/sdk-go/internal/remote/browser/models"
	request2 "github.com/smash-hq/sdk-go/internal/remote/request"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"net/url"
)

func (c *Client) ScrapingBrowserCreate(ctx context.Context, req *models.CreateBrowserRequest) (*models.CreateBrowserResponse, error) {
	value := &url.Values{}
	value.Set("token", req.ApiKey)
	value.Set("proxy_country", req.Proxy.Country)
	value.Set("proxy_url", req.Proxy.Url)
	value.Set("session_id", req.Proxy.SessionId)
	value.Set("session_duration", fmt.Sprintf("%d", req.Proxy.SessionDuration))
	value.Set("gateway", req.Proxy.Gateway)
	value.Set("channel_id", req.Proxy.ChannelId)
	if req.Input != nil {
		for k, v := range req.Input {
			value.Set(k, v)
		}
	}
	parse, _ := url.Parse(fmt.Sprintf("%s/browser", c.BaseUrl))
	parse.RawQuery = value.Encode()
	request, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodGet,
		Url:    parse.String(),
	})
	if err != nil {
		return nil, err
	}

	var task *models.CreateBrowserResponse
	err = json.Unmarshal([]byte(request), &task)
	if err != nil {
		return nil, status.Error(codes.Internal, "create task failed, unmarshal response body error")
	}
	if !task.Success {
		return nil, status.Errorf(codes.Internal, "create task failed, code: %d, message: %s", task.Code, task.Message)
	}

	u, err := url.Parse(c.BaseUrl)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "parse url error: %s", err.Error())
	}
	devValue := &url.Values{}
	devValue.Set("token", req.ApiKey)
	if req.Input != nil {
		for k, v := range req.Input {
			devValue.Set(k, v)
		}
	}
	if req.Proxy.Country != "" {
		devValue.Set("proxy_country", req.Proxy.Country)
	}
	task.DevtoolsUrl = fmt.Sprintf("wss://%s/browser/%s?%s", u.Host, task.TaskId, devValue.Encode())
	return task, nil
}
