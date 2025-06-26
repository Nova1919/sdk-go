package storage_http

import (
	"context"
	"encoding/json"
	"fmt"
	request2 "github.com/scrapeless-ai/sdk-go/internal/remote/request"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage/models"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"net/http"
)

func (c *Client) ListDatasets(ctx context.Context, req *models.ListDatasetsRequest) (*models.ListDatasetsResponse, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodGet,
		Url:     fmt.Sprintf("%s/api/v1/dataset?actorId=%s&desc=%v&page=%d&pageSize=%d&runId=%s", c.BaseUrl, *req.ActorId, req.Desc, req.Page, req.PageSize, *req.RunId),
		Body:    "",
		Headers: map[string]string{},
	})
	log.Infof("list dataset body:%s", body)
	if err != nil {
		log.Errorf("list dataset err:%v", err)
		return nil, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	if resp.Err {
		return nil, fmt.Errorf("get dataset list err:%s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	var respData models.ListDatasetsResponse
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	return &respData, nil
}

func (c *Client) CreateDataset(ctx context.Context, req *models.CreateDatasetRequest) (*models.Dataset, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	body, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodPost,
		Url:     fmt.Sprintf("%s/api/v1/dataset", c.BaseUrl),
		Body:    string(reqBody),
		Headers: map[string]string{},
	})
	log.Infof("create dataset body:%s", body)
	if err != nil {
		log.Errorf("create dataset err:%v", err)
		return nil, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	if resp.Err {
		return nil, fmt.Errorf("creat dataset err:%s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	var respData models.Dataset
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	return &respData, nil
}

func (c *Client) UpdateDataset(ctx context.Context, datasetID, name string) (bool, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodPut,
		Url:     fmt.Sprintf("%s/api/v1/dataset/%s", c.BaseUrl, datasetID),
		Body:    fmt.Sprintf(`{"name":"%s"}`, name),
		Headers: map[string]string{},
	})
	log.Infof("up dataset body:%s", body)
	if err != nil {
		log.Errorf("up dataset err:%v", err)
		return false, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return false, err
	}
	if resp.Err {
		return false, fmt.Errorf("edit dataset err:%s", resp.Msg)
	}
	return true, nil
}
func (c *Client) DelDataset(ctx context.Context, datasetID string) (bool, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodDelete,
		Url:     fmt.Sprintf("%s/api/v1/dataset/%s", c.BaseUrl, datasetID),
		Headers: map[string]string{},
	})
	log.Infof("del dataset body:%s", body)
	if err != nil {
		log.Errorf("del dataset err:%v", err)
		return false, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return false, err
	}
	if resp.Err {
		return false, fmt.Errorf("del dataset err:%s", resp.Msg)
	}
	return true, nil
}

func (c *Client) GetDataset(ctx context.Context, req *models.GetDataset) (*models.DatasetItem, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodGet,
		Url:     fmt.Sprintf("%s/api/v1/dataset/%s/items?page=%d&pageSize=%d&desc=%v", c.BaseUrl, req.DatasetId, req.Page, req.PageSize, req.Desc),
		Headers: map[string]string{},
	})
	log.Infof("get dataset body:%s", body)
	if err != nil {
		log.Errorf("get dataset err:%v", err)
		return nil, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	if resp.Err {
		return nil, fmt.Errorf("get dataset item err:%s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	var respData models.DatasetItem
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	return &respData, nil
}

func (c *Client) AddDatasetItem(ctx context.Context, datasetId string, data []map[string]any) (bool, error) {
	reqBody, err := json.Marshal(map[string]any{
		"items": data,
	})
	if err != nil {
		log.Errorf("marshal dataset item err:%v", err)
		return false, err
	}
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodPost,
		Url:     fmt.Sprintf("%s/api/v1/dataset/%s/items", c.BaseUrl, datasetId),
		Body:    string(reqBody),
		Headers: map[string]string{},
	})
	log.Infof("add dataset item body:%s", body)
	if err != nil {
		log.Errorf("add dataset item err:%v", err)
		return false, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Errorf("unmarshal dataset item err:%v", err)
		return false, err
	}
	if resp.Err {
		log.Errorf("add dataset item err:%s", resp.Msg)
		return false, fmt.Errorf("add dataset item err:%s", resp.Msg)
	}
	return true, nil
}
