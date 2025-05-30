package scraping

type Actor string

var (
	ScraperAmazon Actor = "scraper.amazon"
)

type ScrapingTaskRequest struct {
	// Actor identifier (e.g., "scraper.amazon", "scraper.walmart")
	Actor Actor `json:"actor"`

	// Input parameters for the scraper
	Input map[string]interface{} `json:"input"`

	// Proxy configuration
	ProxyCountry string `json:"country"`
}

type ScrapingTaskResponse struct {
	// Response message
	Message string `json:"message"`

	// Task ID for tracking the request
	TaskID string `json:"taskId"`
}
