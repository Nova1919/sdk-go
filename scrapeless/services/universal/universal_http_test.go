package universal

import (
	"context"
	"testing"
)

func TestUniversalHttp_CreateTask(t *testing.T) {
	universal := New()
	task, err := universal.CreateTask(context.Background(), UniversalTaskRequest{
		Actor: ScraperUniversal,
		Input: map[string]any{
			"url":       "https://www.google.com/",
			"js_render": true,
		},
		ProxyCountry: "US",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(task))
}
