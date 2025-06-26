package http

import (
	"fmt"
	"github.com/smash-hq/sdk-go/env"
	"github.com/smash-hq/sdk-go/scrapeless/log"
	"io"
	"net/http"
)

func (c *Client) Request(keyword string, method string, path string, body io.Reader, headers map[string]string) (data []byte, err error) {
	if path != "" && path[0] == '/' {
		path = path[1:]
	}
	u := fmt.Sprintf("%s/api/v1/run/%s/%s", c.BaseUrl, keyword, path)
	request, err := http.NewRequest(method, u, body)
	if err != nil {
		log.Errorf("new request error :%v", err)
		return nil, err
	}
	for k, v := range headers {
		request.Header.Set(k, v)
	}
	request.Header.Set(env.Env.HTTPHeader, env.GetActorEnv().ApiKey)
	do, err := c.client.Do(request)
	if err != nil {
		log.Errorf("do request error :%v", err)
		return nil, err
	}
	b, err := io.ReadAll(do.Body)
	if err != nil {
		log.Errorf("read body error :%v", err)
		return nil, err
	}
	log.Infof("request body :%s", string(b))
	defer do.Body.Close()
	return b, nil
}
