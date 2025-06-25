package main

import (
	"context"
	"github.com/smash-hq/sdk-go/scrapeless"
	"github.com/smash-hq/sdk-go/scrapeless/log"
	proxy2 "github.com/smash-hq/sdk-go/scrapeless/services/proxies"
	"io"
	"net/http"
	"net/url"
)

func main() {
	client := scrapeless.New(scrapeless.WithProxy())
	defer client.Close()
	proxy, err := client.Proxy.Proxy(context.TODO(), proxy2.ProxyActor{
		Country:         "US",
		SessionDuration: 180,
		SessionId:       "YOUR SESSION ID",
		Gateway:         "YOU GATEWAY",
	})
	if err != nil {
		panic(err)
	}
	log.Info(proxy)
	parse, err := url.Parse(proxy)
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("HEADER KEY", "HEADER VALUE")
	clientHttp := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(parse)}}
	resp, err := clientHttp.Do(request)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Info(string(body))
}
