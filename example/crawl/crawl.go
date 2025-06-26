package main

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/scrapeless"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/crawl"
)

func main() {
	client := scrapeless.New(scrapeless.WithCrawl())

	// Crawl
	response, err := client.Crawl.CrawlUrl(context.Background(), "https://redditinc.com/blog", crawl.CrawlParams{
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
		panic(err)
	}
	log.Infof("Crawl response: %v", response)

	// scrape
	scrapeResponse, err := client.Crawl.ScrapeUrl(context.Background(), "https://docs.scrapeless.com/en/overview/", crawl.ScrapeOptions{
		BrowserOptions: crawl.ICreateBrowser{
			SessionName:      "Crawl",
			SessionTTL:       "900",
			SessionRecording: "true",
			ProxyCountry:     "ANY",
		},
	})
	if err != nil {
		panic(err)
	}
	log.Infof("Scrape response: %v", scrapeResponse)
}
