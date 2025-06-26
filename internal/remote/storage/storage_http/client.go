package storage_http

import (
	"github.com/scrapeless-ai/sdk-go/env"
	"google.golang.org/grpc"
	"net/http"
)

var (
	defaultStorageClient *Client
)

type Client struct {
	client  *http.Client
	BaseUrl string
}

func Init(baseUrl ...string) {
	var (
		err error
	)
	u := env.Env.ScrapelessStorageUrl
	if len(baseUrl) > 0 {
		u = baseUrl[0]
	}
	defaultStorageClient, err = New(u)
	defaultStorageClient.regisHttpHandleFunc()
	if err != nil {
		panic(err)
	}
}

func Default() *Client {
	return defaultStorageClient
}

func New(baseUrl string, opts ...grpc.DialOption) (*Client, error) {
	return &Client{
		client:  &http.Client{},
		BaseUrl: baseUrl,
	}, nil
}

func (c *Client) Close() error {
	return nil
}
