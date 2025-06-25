package deepserp

type Actor string

var (
	ScraperAmazon Actor = "scraper.amazon"
)

type DeepserpTaskRequest struct {
	// Actor identifier (e.g., "scraper.google.search")
	Actor Actor `json:"actor"`

	// Input parameters for the deepserp
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
