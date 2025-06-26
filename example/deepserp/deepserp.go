package main

import (
	"context"
	scrapeless "github.com/scrapeless-ai/sdk-go/scrapeless"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/deepserp"
)

func main() {
	client := scrapeless.New(scrapeless.WithDeepSerp())

	scrape, err := client.DeepSerp.Scrape(context.Background(), deepserp.DeepserpTaskRequest{
		Actor: "scraper.google.search",
		Input: map[string]interface{}{
			"q": "nike site:www.nike.com",
		},
		ProxyCountry: "US",
	})
	if err != nil {
		log.Errorf("scraping create err:%v", err)
		return
	}
	log.Infof("%+v", scrape)
}
