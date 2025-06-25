package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/smash-hq/sdk-go/internal/remote/crawl/models"
	request2 "github.com/smash-hq/sdk-go/internal/remote/request"
	"github.com/tidwall/gjson"
	"net/http"
)

// ScrapeUrl scrape a single url
func (c *Client) ScrapeUrl(ctx context.Context, req *models.ScrapeOptions) (id string, err error) {
	body, _ := json.Marshal(req)
	response, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodPost,
		Url:     fmt.Sprintf("%s/api/v1/crawler/scrape", c.BaseUrl),
		Body:    string(body),
		Headers: map[string]string{},
	})
	if err != nil {
		return "", err
	}
	data := gjson.Parse(response)
	success := data.Get("success").Bool()
	if !success {
		return "", errors.New("server error")
	}
	return data.Get("id").String(), nil
}

func (c *Client) BatchScrapeUrls(ctx context.Context, req *models.ScrapeOptionsMultiple) (scrapeResponse *models.ScrapeResponse, err error) {
	body, _ := json.Marshal(req)
	response, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodPost,
		Url:     fmt.Sprintf("%s/api/v1/crawler/scrape/batch", c.BaseUrl),
		Body:    string(body),
		Headers: map[string]string{},
	})
	if err != nil {
		return nil, err
	}
	data := gjson.Parse(response)
	success := data.Get("success").Bool()
	if !success {
		return nil, errors.New("server error")
	}
	var invalidURLs = make([]string, 0)
	id := data.Get("id").String()
	array := data.Get("invalidURLs").Array()
	for _, result := range array {
		invalidURLs = append(invalidURLs, result.String())
	}
	return &models.ScrapeResponse{
		ID:          id,
		InvalidURLs: invalidURLs,
	}, nil
}

func (c *Client) CheckScrapeStatus(ctx context.Context, id string) (scrapeStatusResponse *models.ScrapeStatusResponse, err error) {
	response, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodGet,
		Url:     fmt.Sprintf("%s/api/v1/crawler/scrape/%s", c.BaseUrl, id),
		Headers: map[string]string{},
	})
	if err != nil {
		return nil, err
	}
	data := gjson.Parse(response)
	success := data.Get("success").Bool()
	if !success {
		return nil, errors.New("server error")
	}
	if err = json.Unmarshal([]byte(response), &scrapeStatusResponse); err != nil {
		return nil, err
	}
	if scrapeStatusResponse.Success {
		return scrapeStatusResponse, nil
	}
	return nil, errors.New(scrapeStatusResponse.Error)
}

func (c *Client) CheckBatchScrapeStatus(ctx context.Context, id string) (scrapeStatusResponseMultiple *models.ScrapeStatusResponseMultiple, err error) {
	response, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodGet,
		Url:     fmt.Sprintf("%s/api/v1/crawler/scrape/batch/%s", c.BaseUrl, id),
		Headers: map[string]string{},
	})
	if err != nil {
		return nil, err
	}
	data := gjson.Parse(response)
	success := data.Get("success").Bool()
	if !success {
		return nil, errors.New("server error")
	}
	if err = json.Unmarshal([]byte(response), &scrapeStatusResponseMultiple); err != nil {
		return nil, err
	}
	return
}

func (c *Client) CrawlUrl(ctx context.Context, req *models.CrawlParams) (id string, err error) {
	body, _ := json.Marshal(req)
	response, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodPost,
		Url:     fmt.Sprintf("%s/api/v1/crawler/crawl", c.BaseUrl),
		Body:    string(body),
		Headers: map[string]string{},
	})
	if err != nil {
		return "", err
	}
	success := gjson.Parse(response).Get("success").Bool()
	if !success {
		return "", errors.New("server error")
	}
	return gjson.Parse(response).Get("id").String(), nil
}

func (c *Client) CheckCrawlStatus(ctx context.Context, id string) (crawlStatusResponse *models.CrawlStatusResponse, err error) {
	response, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodGet,
		Url:     fmt.Sprintf("%s/api/v1/crawler/crawl/%s", c.BaseUrl, id),
		Headers: map[string]string{},
	})
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal([]byte(response), &crawlStatusResponse); err != nil {
		return nil, err
	}
	return
}

func (c *Client) CheckCrawlErrors(ctx context.Context, id string) (crawlErrorsResponse *models.CrawlErrorsResponse, err error) {
	response, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodGet,
		Url:     fmt.Sprintf("%s/api/v1/crawler/crawl/%s/errors", c.BaseUrl, id),
		Headers: map[string]string{},
	})
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal([]byte(response), &crawlErrorsResponse); err != nil {
		return nil, err
	}
	return
}

func (c *Client) CancelCrawl(ctx context.Context, id string) (errorResponse *models.ErrorResponse, err error) {
	response, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodDelete,
		Url:     fmt.Sprintf("%s/api/v1/crawler/crawl/%s", c.BaseUrl, id),
		Headers: map[string]string{},
	})
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal([]byte(response), &errorResponse); err != nil {
		return nil, err
	}
	if errorResponse.Error != "" {
		return nil, errors.New(errorResponse.Error)
	}
	return errorResponse, nil
}
