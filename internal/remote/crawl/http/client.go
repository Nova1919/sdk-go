package http

import (
	log "github.com/sirupsen/logrus"
	"github.com/smash-hq/sdk-go/env"
	"net/http"
)

var (
	defaultCrawlClient *Client
)

func Init(baseUrl ...string) {
	log.Info("crawl init")
	var err error
	u := env.Env.ScrapelessBaseApiUrl
	if len(baseUrl) > 0 {
		u = baseUrl[0]
	}
	defaultCrawlClient, err = New(u)
	if err != nil {
		panic(err)
	}
}

type Client struct {
	client  *http.Client
	BaseUrl string
}

func Default() *Client {
	return defaultCrawlClient
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
