package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/scrapeless-ai/sdk-go/internal/remote/actor/models"
	request2 "github.com/scrapeless-ai/sdk-go/internal/remote/request"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"github.com/tidwall/gjson"
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) Run(ctx context.Context, req *models.IRunActorData) (string, error) {
	reqBody, _ := json.Marshal(req)
	body, err := request2.RequestData(ctx, request2.ReqInfo{
		Method:  http.MethodPost,
		Url:     fmt.Sprintf("%s/api/v1/actors/%s/runs", c.BaseUrl, req.ActorId),
		Body:    string(reqBody),
		Headers: map[string]string{},
	})
	if err != nil {
		return "", err
	}
	runId := gjson.Parse(string(body)).Get("runId").String()
	return runId, nil
}

func (c *Client) GetRunInfo(ctx context.Context, runId string) (*models.RunInfo, error) {
	body, err := request2.RequestData(ctx, request2.ReqInfo{
		Method:  http.MethodGet,
		Url:     fmt.Sprintf("%s/api/v1/actors/runs/%s", c.BaseUrl, runId),
		Body:    "",
		Headers: map[string]string{},
	})
	if err != nil {
		log.Errorf("get runInfo err:%v", err)
		return nil, err
	}
	var respData models.RunInfo
	err = json.Unmarshal(body, &respData)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	return &respData, nil
}

func (c *Client) AbortRun(ctx context.Context, actorId, runId string) (bool, error) {
	_, err := request2.RequestData(ctx, request2.ReqInfo{
		Method:  http.MethodDelete,
		Url:     fmt.Sprintf("%s/api/v1/actors/%s/runs/%s", c.BaseUrl, actorId, runId),
		Headers: map[string]string{},
	})
	if err != nil {
		log.Errorf("abort run err:%v", err)
		return false, err
	}
	return true, nil
}

func (c *Client) Build(ctx context.Context, actorId string, version string) (string, error) {
	fmt.Println(fmt.Sprintf("%s/api/v1/actors/%s/builds", c.BaseUrl, actorId))
	body, err := request2.RequestData(ctx, request2.ReqInfo{
		Method:  http.MethodPost,
		Url:     fmt.Sprintf("%s/api/v1/actors/%s/builds", c.BaseUrl, actorId),
		Body:    fmt.Sprintf(`{"version": "%s"}`, version),
		Headers: map[string]string{},
	})
	if err != nil {
		log.Errorf("build err:%v", err)
		return "", err
	}
	buildId := gjson.Parse(string(body)).Get("buildId").String()
	return buildId, nil
}

func (c *Client) GetBuildStatus(ctx context.Context, actorId string, buildId string) (*models.BuildInfo, error) {
	fmt.Println(fmt.Sprintf("%s/api/v1/actors/%s/builds/%s", c.BaseUrl, actorId, buildId))
	body, err := request2.RequestData(ctx, request2.ReqInfo{
		Method:  http.MethodGet,
		Url:     fmt.Sprintf("%s/api/v1/actors/%s/builds/%s", c.BaseUrl, actorId, buildId),
		Headers: map[string]string{},
	})
	if err != nil {
		log.Errorf("get build status err:%v", err)
		return nil, err
	}
	var respData models.BuildInfo
	err = json.Unmarshal(body, &respData)
	return &respData, nil
}

func (c *Client) AbortBuild(ctx context.Context, actorId string, buildId string) (bool, error) {
	body, err := request2.RequestData(ctx, request2.ReqInfo{
		Method:  http.MethodDelete,
		Url:     fmt.Sprintf("%s/api/v1/actors/%s/builds/%s", c.BaseUrl, actorId, buildId),
		Headers: map[string]string{},
	})
	if err != nil {
		log.Errorf("abort build err:%v", err)
		return false, err
	}
	success := gjson.Parse(string(body)).Get("success").Bool()
	return success, nil
}

func (c *Client) GetRunList(ctx context.Context, paginationParams *models.IPaginationParams) ([]models.Payload, error) {
	parse, err := url.Parse(fmt.Sprintf("%s/api/v1/actors/runs", c.BaseUrl))
	if err != nil {
		log.Errorf("parse url err:%v", err)
		return nil, err
	}
	val := &url.Values{}
	val.Set("page", fmt.Sprintf("%d", paginationParams.Page))
	val.Set("pageSize", fmt.Sprintf("%d", paginationParams.PageSize))
	val.Set("desc", strconv.FormatBool(paginationParams.Desc))
	parse.RawQuery = val.Encode()
	body, err := request2.RequestData(ctx, request2.ReqInfo{
		Method:  http.MethodGet,
		Url:     parse.String(),
		Headers: map[string]string{},
	})
	if err != nil {
		log.Errorf("get run list err:%v", err)
		return nil, err
	}
	items := gjson.Parse(string(body)).Get("items").String()
	var respData = []models.Payload{}
	err = json.Unmarshal([]byte(items), &respData)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	return respData, nil
}
