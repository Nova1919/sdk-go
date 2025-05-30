package scraping

// ScrapingRequest defines parameters for a scraping operation
type ScrapingRequest struct {
	// Target website identifier (e.g., "amazon", "walmart", "ebay")
	Site string `json:"site"`

	// URL to scrape
	URL string `json:"url"`

	// Scraping operation timeout (milliseconds)
	Timeout *int `json:"timeout,omitempty"`

	// Whether to execute JavaScript (for SPA websites)
	RenderJS *bool `json:"renderJs,omitempty"`

	// Custom request headers
	Headers map[string]string `json:"headers,omitempty"`

	// Proxy configuration
	Proxy *ProxyConfig `json:"proxy,omitempty"`

	// Region settings, affecting the content displayed by the website
	Geo *GeoSettings `json:"geo,omitempty"`

	// Custom scraping options
	Options map[string]interface{} `json:"options,omitempty"`
}

// ProxyConfig represents proxy settings for a scraping request
type ProxyConfig struct {
	// Proxy type: http, https, socks4, socks5
	Type string `json:"type"`

	// Proxy host
	Host string `json:"host"`

	// Proxy port
	Port int `json:"port"`

	// Proxy username (optional)
	Username *string `json:"username,omitempty"`

	// Proxy password (optional)
	Password *string `json:"password,omitempty"`
}

// GeoSettings represents geographic region settings
type GeoSettings struct {
	// Country code (e.g., "us", "uk")
	Country string `json:"country"`

	// Region/state (optional)
	Region *string `json:"region,omitempty"`

	// City (optional)
	City *string `json:"city,omitempty"`
}

type ScrapingResult[T any] struct {
	// Scraping status: "success" or "error"
	Status string `json:"status"`

	// Scraping result data (optional)
	Data T `json:"data,omitempty"`

	// Error message if the scraping failed
	Error *string `json:"error,omitempty"`

	// Request ID, used for tracking and debugging
	RequestID string `json:"requestId"`

	// Scraping elapsed time in milliseconds
	ElapsedTime int `json:"elapsedTime"`

	// Scraping timestamp (RFC3339 or ISO 8601)
	Timestamp string `json:"timestamp"`
}

type ScrapingTaskRequest struct {
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

type ScrapingTaskResponse struct {
	// Response message
	Message string `json:"message"`

	// Task ID for tracking the request
	TaskID string `json:"taskId"`
}
