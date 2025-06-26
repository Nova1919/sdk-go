package http

import (
	"context"
	"encoding/json"
	"fmt"
	request2 "github.com/smash-hq/sdk-go/internal/remote/request"
	"github.com/smash-hq/sdk-go/internal/remote/scraping/models"
	"github.com/smash-hq/sdk-go/scrapeless/log"
	"net/http"
)

func (c *Client) Scrape(ctx context.Context, req *models.ScrapingRequest) ([]byte, error) {
	body, _ := json.Marshal(req)
	response, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodPost,
		Url:     fmt.Sprintf("%s/scraping", c.BaseUrl),
		Body:    string(body),
		Headers: map[string]string{},
	})
	if err != nil {
		return nil, err
	}
	return []byte(response), nil
}

func (c *Client) CreateTask(ctx context.Context, req *models.ScrapingTaskRequest) ([]byte, error) {
	body, _ := json.Marshal(req)
	response, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodPost,
		Url:     fmt.Sprintf("%s/api/v1/scraper/request", c.BaseUrl),
		Body:    string(body),
		Headers: map[string]string{},
	})
	if err != nil {
		return nil, err
	}
	return []byte(response), nil
}

func (c *Client) GetTaskResult(ctx context.Context, taskIKd string) ([]byte, error) {
	response, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodGet,
		Url:     fmt.Sprintf("%s/api/v1/result/%s", c.BaseUrl, taskIKd),
		Headers: map[string]string{},
	})
	log.Infof("get task result:%s", response)
	if err != nil {
		log.Errorf("get task result err:%v", err)
		return nil, err
	}
	return []byte(response), nil
}
