package main

import (
	"context"
	"github.com/smash-hq/sdk-go/scrapeless"
	"github.com/smash-hq/sdk-go/scrapeless/log"
	"github.com/smash-hq/sdk-go/scrapeless/services/universal"
)

func main() {
	client := scrapeless.New(scrapeless.WithUniversal())
	task, err := client.Universal.CreateTask(context.Background(), universal.UniversalTaskRequest{
		Actor: universal.ScraperUniversal,
		Input: map[string]any{
			"url":       "https://www.scrapeless.com",
			"headless":  false,
			"js_render": true,
			"js_instructions": []any{
				map[string]any{"wait": 10000},
				map[string]any{"wait_for": []any{".dynamic-content", 30000}},
				map[string]any{"click": []any{"#load-more", 1000}},
				map[string]any{"fill": []any{"#search-input", "search term"}},
				map[string]any{"keyboard": []any{"press", "Enter"}},
				map[string]any{"evaluate": "window.scrollTo(0, document.body.scrollHeight)"},
			},
			"block": map[string]any{
				"resources": []string{
					"stylesheet", "image", "media", "font", "script", "texttrack",
					"xhr", "fetch", "eventsource", "websocket", "manifest", "other",
				},
			},
		},
	})
	if err != nil {
		log.Errorf("scraping create err:%v", err)
		return
	}
	log.Infof("%+v", task)
	// webUnlocker
	webUnlocker, err := client.Universal.Scrape(context.Background(), universal.UniversalTaskRequest{
		Actor: universal.ScraperUniversal,
		Input: map[string]any{
			"url":      "https://www.nike.com/ca/launch?s=upcoming",
			"type":     "",
			"redirect": false,
			"method":   "GET",
			"header": map[string]any{
				"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) " +
					"AppleWebKit/537.36 (KHTML, like Gecko) " +
					"Chrome/130.0.0.0 Safari/537.36",
			},
		},
	})
	if err != nil {
		log.Errorf("scraping create err:%v", err)
		return
	}
	log.Infof("%+v", webUnlocker)

	// akamaiWebSensor
	akamaiWebSensor, err := client.Universal.Scrape(context.Background(), universal.UniversalTaskRequest{
		Actor: universal.ScraperAkamaiwebUniversal,
		Input: map[string]any{
			"abck": "0DF99F296DFB4060D36BB36ED3D54130~-1~YAAQ07khF8hvCq+SAQAA54Wiwgwlmf8ZB5T97F80Saj/q5/FwHj9hAOqCdcV+cKqxRtZZ0UK6+bALBxlnay+3dlZcAWfnN+UmDoKmRO9kn+wGN7Z62Oob3T2o9MWsul10NsFVgNba2pRGhB0IF4fzptFVJBpaekcKsmzcDiquqFbC52KdsSlp21CG42NisM6dOXlx2gVSeqLU2ijNVpMaEZeOA0ItvLauaf/EGXWumZZVwvXGcEOFBXrWV63NA/hp41G0+cisUbClFfHdQAwsGWZyT0xPrus4J3ooZKm5pwVkKKM3kW0+Gu/4xWlEL7N2yGpbYcCUt5pJKRdlIYBF/k33A0dPFbBiwqwJnmeCg/sudUBX0ZuTx2awTc7oEAUVBYJrRsruRhub3iBpGRrnCfe9N/gS21Hc3mTZvo2og2CAMmaVmix5fKJSlQICqNaMLlE9ilg3JqC6k9zePCODtzDwybvA5SX8Y+0Ykf1mwVxDsAeaIQVJiJUVn1JXi9Owt7Qn4+GeJyTv4Ou7gdNrLBGFt0IswnzvEHOsWQiwakEyIJ9K5udFE6qQIjACBRFHNb2feKwSOXpHeXOcrOGO38+bHs+JaY2iN4y1wrJVLq/hlm5vGIfo20NJHNzWsRfmn502DlJorpAR5HfNHbghh0gVJP6tYUvyi6WsMFBMGWfVFgLQs6zyWe9V7U4UnJfH6ImIjSr4PonXqshGNSxMvQUMeg1NQAt6emgVG/0lpsaUvYVwWGQYqTJdJBDsvFjQJWRrfcPvDdwiUNAPJZ0WELf5X7q+eWC7C6lOaMvzME3i00IlwdIzMcRPW1yqrFEMPwN9sTA30aQ7LM05uiYY1aJ+ECi9EOeUgz3fMq5tL8pOuMSWU5dHVuBckGfBOrL6aVMdAtUHfskIMKXVBguag==~-1~||0||~-1",
			"bmsz": "3A6B16B3276CFCE254BEFBEE3CFA915F~YAAQ07khF/NuCq+SAQAAoXqiwhkZpSRVMT7qPsbrnuBKL5cMCX6rHqdhOz/KLauSYLXxO/SlnflKmmQp0W0Q8uZ6UNC6SRXYctgtcpSTV32rCgnozZGFlk1/nVGRAjMY1UH+97v7EhDOSynkGPqXmxRF3fNd0pd31Fu+ZddDU/U/8PiMrrDqEjt4aaibL+GGbjn+rXR8cL+XrJn2JzIbZm+ddrvYNX+K2Q8WchPywJyLWx7Owq0uRB4OOzqWKz5TZD3dWlirj09fI1YUlnBRlL/mVnSXCWwekVucauI652lV6wqfs1Iykm4ifZkxuuUc9IcaLTpby6clHTZ0UH50eN9WMyiRfT5SNnkqZkVjVg08oYYRXAUdnNWPHqj849yTQdnH1KynNRM5B0U45ApGMnLRUOixTsXAFHU=~3682628~3160113",
			"url":  "https://www.scrapeless.com/",
			"userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) " +
				"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36",
		},
	})
	if err != nil {
		log.Errorf("scraping create err:%v", err)
		return
	}
	log.Infof("%+v", akamaiWebSensor)
}
