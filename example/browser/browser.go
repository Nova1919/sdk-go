package main

import (
	"context"
	"github.com/smash-hq/sdk-go/scrapeless"
	"github.com/smash-hq/sdk-go/scrapeless/log"
	"github.com/smash-hq/sdk-go/scrapeless/services/browser"
)

func main() {
	client := scrapeless.New(scrapeless.WithBrowser())
	defer client.Close()

	browserInfo, err := client.Browser.Create(context.Background(), browser.Actor{
		Input:        browser.Input{SessionTtl: "180"},
		ProxyCountry: "US",
	})
	if err != nil {
		panic(err)
	}
	log.Infof("%+v", browserInfo)
}
