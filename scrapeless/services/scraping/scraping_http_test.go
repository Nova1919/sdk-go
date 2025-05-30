package scraping

import (
	"context"
	"testing"
)

func TestScrapingHttp_CreateTask(t *testing.T) {
	scraping := New()
	task, err := scraping.CreateTask(context.Background(), ScrapingTaskRequest{
		Actor: "scraper.tiktok.mobile.shop.detail",
		Input: map[string]any{
			"region":     "VN",
			"product_id": "1729443471139309751",
			"locale":     "",
		},
		ProxyCountry: "US",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(string(task))
}
