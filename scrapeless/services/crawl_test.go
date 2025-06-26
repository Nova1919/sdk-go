package services

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/crawl"
	"testing"
)

func TestCrawl(t *testing.T) {
	a := crawl.New()
	id, err := a.CrawlUrl(context.Background(), "https://redditinc.com/blog", crawl.CrawlParams{
		Limit: 10,
		ScrapeOptions: crawl.ScrapeOptions{
			Formats: []string{"links",
				"markdown",
				"html",
				"screenshot"},
		},
		BrowserOptions: crawl.ICreateBrowser{
			SessionName:      "Crawl",
			SessionTTL:       "900",
			SessionRecording: "true",
			ProxyCountry:     "ANY",
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", id)

	//response, err := a.CheckCrawlStatus(context.Background(), "38cb4977-9e75-4f6c-b193-e601c5073c91")
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Logf("%+v\n", response)

	//errorsResponse, err := a.CheckCrawlErrors(context.Background(), "666d3881-94c2-4383-9d5b-622d5d07e597")
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Log(errorsResponse)
}

func TestScrape(t *testing.T) {
	a := crawl.New()
	id, err := a.AsyncScrapeUrl(context.Background(), "https://docs.scrapeless.com/en/overview/", crawl.ScrapeOptions{
		BrowserOptions: crawl.ICreateBrowser{
			SessionName:      "Crawl",
			SessionTTL:       "900",
			SessionRecording: "true",
			ProxyCountry:     "ANY",
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", id)
	//response, err := a.CheckScrapeStatus(context.Background(), "0c10c021-f0d8-4994-832d-edbbd1d29f48")
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Logf("%+v\n", response)
}
