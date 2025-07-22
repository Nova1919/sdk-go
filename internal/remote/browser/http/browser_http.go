package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/scrapeless-ai/sdk-go/internal/remote/browser/models"
	request2 "github.com/scrapeless-ai/sdk-go/internal/remote/request"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"net/url"
)

func (c *Client) ScrapingBrowserCreate(ctx context.Context, req *models.CreateBrowserRequest) (*models.CreateBrowserResponse, error) {
	fingerprint, _ := json.Marshal(req.Fingerprint)
	value := &url.Values{}
	value.Set("token", req.ApiKey)
	value.Set("session_name", req.SessionName)
	value.Set("session_ttl", fmt.Sprintf("%v", req.SessionTtl))
	value.Set("session_recording", fmt.Sprintf("%v", req.SessionRecording))
	value.Set("proxy_country", req.ProxyCountry)
	value.Set("proxy_url", req.ProxyUrl)
	value.Set("fingerprint", string(fingerprint))
	value.Set("extension_ids", req.ExtensionIds)
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
	task.DevtoolsUrl = fmt.Sprintf("wss://%s/browser/%s?%s", u.Host, task.TaskId, devValue.Encode())
	return task, nil
}
