package crawl

type ScrapeOptions struct {
	Formats         []string          `json:"formats,omitempty"`
	Headers         map[string]string `json:"headers,omitempty"`
	IncludeTags     []string          `json:"includeTags,omitempty"`
	ExcludeTags     []string          `json:"excludeTags,omitempty"`
	OnlyMainContent bool              `json:"onlyMainContent,omitempty"`
	WaitFor         int               `json:"waitFor,omitempty"`
	Timeout         int               `json:"timeout,omitempty"`
	BrowserOptions  ICreateBrowser    `json:"browserOptions"`
}

type ScrapeOptionsMultiple struct {
	Url             []string          `json:"url"`
	Formats         []string          `json:"formats,omitempty"`
	Headers         map[string]string `json:"headers,omitempty"`
	IncludeTags     []string          `json:"includeTags,omitempty"`
	ExcludeTags     []string          `json:"excludeTags,omitempty"`
	OnlyMainContent bool              `json:"onlyMainContent,omitempty"`
	WaitFor         int               `json:"waitFor,omitempty"`
	Timeout         int               `json:"timeout,omitempty"`
	BrowserOptions  ICreateBrowser    `json:"browserOptions"`
}

type ScrapeResponse struct {
	ID string `json:"id"`

	// If ignoreInvalidURLs is true, this is an array containing the invalid URLs that were specified in the request.
	// If there were no invalid URLs, this will be an empty array.
	// If ignoreInvalidURLs is false, this field will be undefined.
	InvalidURLs []string `json:"success"`
}

type Status string

const (
	StatusScraping  Status = "scraping"
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
	StatusCancelled Status = "cancelled"
	StatusActive    Status = "active"
	StatusPaused    Status = "paused"
	StatusPending   Status = "pending"
	StatusQueued    Status = "queued"
	StatusWaiting   Status = "waiting"
)

type ScrapingCrawlDocumentMetadata struct {
	Title             string   `json:"title,omitempty"`
	Description       string   `json:"description,omitempty"`
	Language          string   `json:"language,omitempty"`
	Keywords          string   `json:"keywords,omitempty"`
	Robots            string   `json:"robots,omitempty"`
	OgTitle           string   `json:"ogTitle,omitempty"`
	OgDescription     string   `json:"ogDescription,omitempty"`
	OgURL             string   `json:"ogUrl,omitempty"`
	OgImage           string   `json:"ogImage,omitempty"`
	OgAudio           string   `json:"ogAudio,omitempty"`
	OgDeterminer      string   `json:"ogDeterminer,omitempty"`
	OgLocale          string   `json:"ogLocale,omitempty"`
	OgLocaleAlternate []string `json:"ogLocaleAlternate,omitempty"`
	OgSiteName        string   `json:"ogSiteName,omitempty"`
	OgVideo           string   `json:"ogVideo,omitempty"`
	DCTermsCreated    string   `json:"dctermsCreated,omitempty"`
	DCDateCreated     string   `json:"dcDateCreated,omitempty"`
	DCDate            string   `json:"dcDate,omitempty"`
	DCTermsType       string   `json:"dctermsType,omitempty"`
	DCType            string   `json:"dcType,omitempty"`
	DCTermsAudience   string   `json:"dctermsAudience,omitempty"`
	DCTermsSubject    string   `json:"dctermsSubject,omitempty"`
	DCSubject         string   `json:"dcSubject,omitempty"`
	DCDescription     string   `json:"dcDescription,omitempty"`
	DCTermsKeywords   string   `json:"dctermsKeywords,omitempty"`
	ModifiedTime      string   `json:"modifiedTime,omitempty"`
	PublishedTime     string   `json:"publishedTime,omitempty"`
	ArticleTag        string   `json:"articleTag,omitempty"`
	ArticleSection    string   `json:"articleSection,omitempty"`
	SourceURL         string   `json:"sourceURL,omitempty"`
	StatusCode        int      `json:"statusCode,omitempty"`
	Error             string   `json:"error,omitempty"`

	// 允许额外的动态字段（等价于 TypeScript 中的 [key: string]: any）
	ExtraFields map[string]interface{} `json:"-"`
}

type ScrapingCrawlDocument struct {
	Markdown   string                        `json:"markdown,omitempty"`
	HTML       string                        `json:"html,omitempty"`
	RawHTML    string                        `json:"rawHtml,omitempty"`
	Links      []string                      `json:"links,omitempty"`
	Extract    any                           `json:"extract,omitempty"`
	Screenshot string                        `json:"screenshot,omitempty"`
	Metadata   ScrapingCrawlDocumentMetadata `json:"metadata,omitempty"`
}

type ScrapeStatusResponse struct {
	Status Status                `json:"status,omitempty"`
	Data   ScrapingCrawlDocument `json:"data,omitempty"`
}

type ScrapeStatusResponseMultiple struct {
	Total     int                     `json:"total"`
	Completed int                     `json:"completed"`
	Status    Status                  `json:"status"`
	Data      []ScrapingCrawlDocument `json:"data"`
}

type ICreateBrowser struct {
	SessionName      string `json:"session_name,omitempty"`
	SessionTTL       string `json:"session_ttl,omitempty"`
	SessionRecording string `json:"session_recording,omitempty"`
	ProxyCountry     string `json:"proxy_country,omitempty"`
	ProxyURL         string `json:"proxy_url,omitempty"`
	Fingerprint      string `json:"fingerprint,omitempty"`
}

type ScrapeParams struct {
	ScrapeOptions
	BrowserOptions *ICreateBrowser `json:"browserOptions,omitempty"`
}

type CrawlParams struct {
	IncludePaths           []string       `json:"includePaths,omitempty"`
	ExcludePaths           []string       `json:"excludePaths,omitempty"`
	MaxDepth               int            `json:"maxDepth,omitempty"`
	MaxDiscoveryDepth      int            `json:"maxDiscoveryDepth,omitempty"`
	Limit                  int            `json:"limit,omitempty,omitempty"`
	AllowBackwardLinks     bool           `json:"allowBackwardLinks,omitempty"`
	AllowExternalLinks     bool           `json:"allowExternalLinks,omitempty"`
	IgnoreSitemap          bool           `json:"ignoreSitemap,omitempty"`
	DeduplicateSimilarURLs bool           `json:"deduplicateSimilarURLs,omitempty"`
	IgnoreQueryParameters  bool           `json:"ignoreQueryParameters,omitempty"`
	RegexOnFullURL         bool           `json:"regexOnFullURL,omitempty"`
	Delay                  int            `json:"delay,omitempty"`
	ScrapeOptions          ScrapeOptions  `json:"scrapeOptions,omitempty"`
	BrowserOptions         ICreateBrowser `json:"browserOptions,omitempty"`
}
type CrawlScrapeOptions struct {
	Formats         []string          `json:"formats,omitempty"`
	Headers         map[string]string `json:"headers,omitempty"`
	IncludeTags     []string          `json:"includeTags,omitempty"`
	ExcludeTags     []string          `json:"excludeTags,omitempty"`
	OnlyMainContent bool              `json:"onlyMainContent,omitempty"`
	WaitFor         int               `json:"waitFor,omitempty"`
	Timeout         int               `json:"timeout,omitempty"`
}

type CrawlResponse struct {
	ID      string `json:"id,omitempty"`
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

type CrawlStatusResponse struct {
	Status    Status                  `json:"status"`
	Completed int                     `json:"completed"`
	Total     int                     `json:"total"`
	Data      []ScrapingCrawlDocument `json:"data"`
}

type CrawlErrorDetail struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp,omitempty"`
	Url       string `json:"url"`
	Error     string `json:"error"`
}

type CrawlErrorsResponse struct {
	Errors        []CrawlErrorDetail `json:"errors"`
	RobotsBlocked []string           `json:"robotsBlocked"`
}

type ErrorResponse struct {
	Status string `json:"success"`
	Error  string `json:"error"`
}
