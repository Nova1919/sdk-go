package storage_http

import (
	"context"
	"encoding/json"
	"fmt"
	request2 "github.com/scrapeless-ai/sdk-go/internal/remote/request"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage/models"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"net/http"
	"net/url"
)

func (c *Client) ListCollections(ctx context.Context, req *models.ListCollectionsRequest) (*models.ListCollectionsResponse, error) {
	u, err := url.Parse(fmt.Sprintf("%s/api/v1/vector?page=%d&pageSize=%d&desc=%t", c.BaseUrl, req.Page, req.PageSize, req.Desc))
	if err != nil {
		return nil, err
	}
	q := u.Query()
	if req.RunId != nil && *req.RunId != "" {
		q.Add("runId", *req.RunId)
	}
	if req.ActorId != nil && *req.ActorId != "" {
		q.Add("actorId", *req.ActorId)
	}
	u.RawQuery = q.Encode()
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodGet,
		Url:     u.String(),
		Headers: map[string]string{},
	})
	log.Infof("list collection body:%s", body)
	if err != nil {
		log.Errorf("list collection err:%v", err)
		return nil, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Errorf("unmarshal resp error:%v", err)
		return nil, err
	}
	if resp.Err {
		return nil, fmt.Errorf("get collections list err:%s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	var respData models.ListCollectionsResponse
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	return &respData, nil
}

func (c *Client) CreateCollections(ctx context.Context, req *models.CreateCollectionRequest) (*models.CreateCollectionResponse, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	body, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodPost,
		Url:     fmt.Sprintf("%s/api/v1/vector", c.BaseUrl),
		Body:    string(reqBody),
		Headers: map[string]string{},
	})
	log.Infof("create collection body:%s", body)
	if err != nil {
		log.Errorf("create collection err:%v", err)
		return nil, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return nil, err
	}
	if resp.Err {
		return nil, fmt.Errorf("get collection list err: %s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	var coll models.Collection
	err = json.Unmarshal(marshal, &coll)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	return &models.CreateCollectionResponse{
		Coll: coll,
	}, nil

}

func (c *Client) UpdateCollection(ctx context.Context, req *models.UpdateCollectionRequest) error {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return err
	}
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodPut,
		Url:     fmt.Sprintf("%s/api/v1/vector/%s", c.BaseUrl, req.CollId),
		Body:    string(reqBody),
		Headers: map[string]string{},
	})
	log.Infof("update collection body:%s", body)
	if err != nil {
		log.Errorf("update collection err:%v", err)
		return err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return err
	}
	if resp.Err {
		return fmt.Errorf("edit collection err:%s", resp.Msg)
	}
	return nil
}

func (c *Client) DelCollection(ctx context.Context, collId string) error {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodDelete,
		Url:     fmt.Sprintf("%s/api/v1/vector/%s", c.BaseUrl, collId),
		Headers: map[string]string{},
	})
	log.Infof("up collection body:%s", body)
	if err != nil {
		log.Errorf("up collection err:%v", err)
		return err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return err
	}
	if resp.Err {
		return fmt.Errorf("edit collection err:%s", resp.Msg)
	}
	return nil
}

func (c *Client) GetCollection(ctx context.Context, collId string) (*models.Collection, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodGet,
		Url:     fmt.Sprintf("%s/api/v1/vector/%s", c.BaseUrl, collId),
		Headers: map[string]string{},
	})
	log.Infof("get collection body:%s", body)
	if err != nil {
		log.Errorf("get collection err:%v", err)
		return nil, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	if resp.Err {
		return nil, fmt.Errorf("get collection item err:%s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	var respData models.Collection
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	return &respData, nil
}

func (c *Client) CreateDocs(ctx context.Context, req *models.CreateDocsRequest) (*models.DocOpResponse, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodPost,
		Url:     fmt.Sprintf("%s/api/v1/vector/%s/docs", c.BaseUrl, req.CollId),
		Body:    string(reqBody),
		Headers: map[string]string{},
	})
	log.Infof("get collection body:%s", body)
	if err != nil {
		log.Errorf("get collection err:%v", err)
		return nil, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	if resp.Err {
		return nil, fmt.Errorf("get collection item err:%s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	var respData models.DocOpResponse
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	return &respData, nil
}

func (c *Client) UpdateDocs(ctx context.Context, req *models.UpdateDocsRequest) (*models.DocOpResponse, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodPut,
		Url:     fmt.Sprintf("%s/api/v1/vector/%s/docs", c.BaseUrl, req.CollId),
		Body:    string(reqBody),
		Headers: map[string]string{},
	})
	log.Infof("get collection body:%s", body)
	if err != nil {
		log.Errorf("get collection err:%v", err)
		return nil, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	if resp.Err {
		return nil, fmt.Errorf("get collection item err:%s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	var respData models.DocOpResponse
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	return &respData, nil
}

func (c *Client) UpsertDocs(ctx context.Context, req *models.UpsertVectorDocsParam) (*models.DocOpResponse, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodPost,
		Url:     fmt.Sprintf("%s/api/v1/vector/%s/docs/upsert", c.BaseUrl, req.CollId),
		Body:    string(reqBody),
		Headers: map[string]string{},
	})
	log.Infof("get collection body:%s", body)
	if err != nil {
		log.Errorf("get collection err:%v", err)
		return nil, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	if resp.Err {
		return nil, fmt.Errorf("get collection item err:%s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	var respData models.DocOpResponse
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	return &respData, nil
}

func (c *Client) DelDocs(ctx context.Context, req *models.DeleteDocsRequest) (*models.DocOpResponse, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodDelete,
		Url:     fmt.Sprintf("%s/api/v1/vector/%s/docs", c.BaseUrl, req.CollId),
		Body:    string(reqBody),
		Headers: map[string]string{},
	})
	log.Infof("get collection body:%s", body)
	if err != nil {
		log.Errorf("get collection err:%v", err)
		return nil, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	if resp.Err {
		return nil, fmt.Errorf("get collection item err:%s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	var respData models.DocOpResponse
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	return &respData, nil
}

func (c *Client) QueryDocs(ctx context.Context, req *models.QueryVectorRequest) ([]*models.Doc, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodPost,
		Url:     fmt.Sprintf("%s/api/v1/vector/%s/docs/query", c.BaseUrl, req.CollId),
		Body:    string(reqBody),
		Headers: map[string]string{},
	})
	log.Infof("get collection body:%s", body)
	if err != nil {
		log.Errorf("get collection err:%v", err)
		return nil, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	if resp.Err {
		return nil, fmt.Errorf("get collection item err:%s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	respData := make([]*models.Doc, 0)
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	return respData, nil
}

func (c *Client) QueryDocsByIds(ctx context.Context, req *models.QueryDocsByIdsRequest) (map[string]*models.Doc, error) {
	u, err := url.Parse(fmt.Sprintf("%s/api/v1/vector/%s/docs", c.BaseUrl, req.CollId))
	if err != nil {
		return nil, err
	}

	query := u.Query()
	for i := range req.Ids {
		query.Add("ids", req.Ids[i])
	}
	u.RawQuery = query.Encode()

	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodGet,
		Url:    u.String(),
	})
	log.Infof("get docs body: %s", body)
	if err != nil {
		log.Errorf("get collection err: %v", err)
		return nil, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	if resp.Err {
		return nil, fmt.Errorf("get collection item err:%s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	respData := make(map[string]*models.Doc)
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		log.Errorf("unmarshal resp error :%v", err)
		return nil, err
	}
	return respData, nil
}
