package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scrapeless-ai/sdk-go/internal/remote/profile/models"
	"github.com/scrapeless-ai/sdk-go/internal/remote/request"
	"net/http"
	"net/url"
)

func (c *Client) Create(ctx context.Context, name string) (profile *models.ProfileInfo, err error) {
	profile = new(models.ProfileInfo)
	req := make(map[string]string)
	req["name"] = name
	body, _ := json.Marshal(req)
	response, err := request.Request(ctx, request.ReqInfo{
		Method:  http.MethodPost,
		Url:     fmt.Sprintf("%s/browser/profiles", c.BaseUrl),
		Body:    string(body),
		Headers: map[string]string{},
	})
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal([]byte(response), profile); err != nil {
		return nil, err
	}
	return
}

func (c *Client) Get(ctx context.Context, profileId string) (profile *models.ProfileInfo, err error) {
	profile = new(models.ProfileInfo)
	response, err := request.Request(ctx, request.ReqInfo{
		Method:  http.MethodGet,
		Url:     fmt.Sprintf("%s/browser/profiles/%s", c.BaseUrl, profileId),
		Headers: map[string]string{},
	})
	if err != nil {
		return nil, err
	}
	if response == "" {
		return nil, errors.New("profile not found")
	}
	if err = json.Unmarshal([]byte(response), profile); err != nil {
		return nil, err
	}

	return
}

func (c *Client) List(ctx context.Context, req *models.ListProfileRequest) (resp *models.ListProfileResponse, err error) {
	resp = new(models.ListProfileResponse)

	u, err := url.Parse(fmt.Sprintf("%s/browser/profiles", c.BaseUrl))
	if err != nil {
		return nil, err
	}
	q := u.Query()
	if req.Name != nil && *req.Name != "" {
		q.Add("name", *req.Name)
	}
	q.Add("page", fmt.Sprintf("%d", req.Page))
	q.Add("pageSize", fmt.Sprintf("%d", req.PageSize))
	u.RawQuery = q.Encode()

	response, err := request.Request(ctx, request.ReqInfo{
		Method:  http.MethodGet,
		Url:     u.String(),
		Headers: map[string]string{},
	})
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal([]byte(response), resp); err != nil {
		return nil, err
	}

	return
}

func (c *Client) Delete(ctx context.Context, profileId string) (resp *models.DeleteProfileResponse, err error) {
	resp = new(models.DeleteProfileResponse)
	response, err := request.Request(ctx, request.ReqInfo{
		Method:  http.MethodDelete,
		Url:     fmt.Sprintf("%s/browser/profiles/%s", c.BaseUrl, profileId),
		Body:    "",
		Headers: map[string]string{},
	})
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal([]byte(response), resp); err != nil {
		return nil, err
	}

	return
}

func (c *Client) Update(ctx context.Context, profileId string, name string) (resp *models.UpdateProfileRequest, err error) {
	resp = new(models.UpdateProfileRequest)
	body, _ := json.Marshal(map[string]string{
		"name": name,
	})
	response, err := request.Request(ctx, request.ReqInfo{
		Method:  http.MethodPut,
		Url:     fmt.Sprintf("%s/browser/profiles/%s", c.BaseUrl, profileId),
		Body:    string(body),
		Headers: map[string]string{},
	})
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal([]byte(response), resp); err != nil {
		return nil, err
	}
	return
}
