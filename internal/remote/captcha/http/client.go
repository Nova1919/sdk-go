package http

import (
	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"net/http"
)

var (
	defaultGatewayClient *Client
)

func Init(baseUrl ...string) {
	log.Info("captcha init")
	var err error
	u := env.Env.ScrapelessBaseApiUrl
	if len(baseUrl) > 0 {
		u = baseUrl[0]
	}
	defaultGatewayClient, err = New(u)
	if err != nil {
		panic(err)
	}
}

type Client struct {
	client  *http.Client
	BaseUrl string
}

func Default() *Client {
	return defaultGatewayClient
}

func New(baseUrl string) (*Client, error) {
	return &Client{
		client:  &http.Client{},
		BaseUrl: baseUrl,
	}, nil
}

func (c *Client) Close() error {
	c.client.CloseIdleConnections()
	return nil
}
