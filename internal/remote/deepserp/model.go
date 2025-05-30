package deepserp

type DeepserpTaskRequest struct {
	// Actor identifier (e.g., "scraper.amazon", "scraper.walmart")
	Actor string `json:"actor"`

	// Input parameters for the scraper
	Input map[string]interface{} `json:"input"`

	// Proxy configuration
	Proxy TaskProxy `json:"proxy"`
}

type TaskProxy struct {
	Country string `json:"country"`
}
