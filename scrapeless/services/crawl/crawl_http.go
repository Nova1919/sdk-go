package crawl

import (
	"context"
	"errors"
	"fmt"
	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/internal/remote/crawl"
	"github.com/scrapeless-ai/sdk-go/internal/remote/crawl/models"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"time"
)

type Crawl struct{}

func New() Crawl {
	log.Info("Internal Crawl init")
	crawl.NewClient("http", env.Env.ScrapelessCrawlApiUrl)
	return Crawl{}
}

func (c *Crawl) ScrapeUrl(ctx context.Context, url string, crawlScrapeOptions ScrapeOptions) (scrapeStatusResponse *ScrapeStatusResponse, err error) {
	id, err := c.AsyncScrapeUrl(ctx, url, crawlScrapeOptions)
	if err != nil {
		return nil, err
	}
	for {
		scrapeStatusResponse, err = c.CheckScrapeStatus(ctx, id)
		if err != nil {
			return nil, err
		}
		switch scrapeStatusResponse.Status {
		case StatusCompleted:
			return scrapeStatusResponse, nil
		case StatusActive, StatusPaused, StatusPending, StatusQueued, StatusWaiting, StatusScraping:
			log.Info("Scraping status: ", scrapeStatusResponse.Status)
			time.Sleep(time.Millisecond * 400)
		default:
			return nil, errors.New(fmt.Sprintf("Crawl job failed or was stopped. Status: %s", scrapeStatusResponse.Status))
		}
	}
}

func (c *Crawl) AsyncScrapeUrl(ctx context.Context, url string, crawlScrapeOptions ScrapeOptions) (id string, err error) {
	id, err = crawl.ClientInterface.ScrapeUrl(ctx, &models.ScrapeOptions{
		Url:             url,
		Formats:         crawlScrapeOptions.Formats,
		Headers:         crawlScrapeOptions.Headers,
		IncludeTags:     crawlScrapeOptions.IncludeTags,
		ExcludeTags:     crawlScrapeOptions.ExcludeTags,
		OnlyMainContent: crawlScrapeOptions.OnlyMainContent,
		WaitFor:         crawlScrapeOptions.WaitFor,
		Timeout:         crawlScrapeOptions.Timeout,
		BrowserOptions: models.ICreateBrowser{
			SessionName:      crawlScrapeOptions.BrowserOptions.SessionName,
			SessionTTL:       crawlScrapeOptions.BrowserOptions.SessionTTL,
			SessionRecording: crawlScrapeOptions.BrowserOptions.SessionRecording,
			ProxyCountry:     crawlScrapeOptions.BrowserOptions.ProxyCountry,
			ProxyURL:         crawlScrapeOptions.BrowserOptions.ProxyURL,
			Fingerprint:      crawlScrapeOptions.BrowserOptions.Fingerprint,
		},
	})
	return
}
func (c *Crawl) CheckScrapeStatus(ctx context.Context, id string) (scrapeStatusResponse *ScrapeStatusResponse, err error) {
	response, err := crawl.ClientInterface.CheckScrapeStatus(ctx, id)
	if err != nil {
		return nil, err
	}
	return &ScrapeStatusResponse{
		Status: Status(response.Status),
		Data:   *c.internalScrapingCrawlDocumentFormat(&response.Data),
	}, nil
}

func (c *Crawl) internalScrapingCrawlDocumentFormat(in *models.ScrapingCrawlDocument) *ScrapingCrawlDocument {
	return &ScrapingCrawlDocument{
		Markdown:   in.Markdown,
		HTML:       in.HTML,
		RawHTML:    in.RawHTML,
		Links:      in.Links,
		Extract:    in.Extract,
		Screenshot: in.Screenshot,
		Metadata: ScrapingCrawlDocumentMetadata{
			Title:             in.Metadata.Title,
			Description:       in.Metadata.DCDescription,
			Language:          in.Metadata.Language,
			Keywords:          in.Metadata.Keywords,
			Robots:            in.Metadata.Robots,
			OgTitle:           in.Metadata.OgTitle,
			OgDescription:     in.Metadata.OgDescription,
			OgURL:             in.Metadata.OgURL,
			OgImage:           in.Metadata.OgImage,
			OgAudio:           in.Metadata.OgAudio,
			OgDeterminer:      in.Metadata.OgDeterminer,
			OgLocale:          in.Metadata.OgLocale,
			OgLocaleAlternate: in.Metadata.OgLocaleAlternate,
			OgSiteName:        in.Metadata.OgSiteName,
			OgVideo:           in.Metadata.OgVideo,
			DCTermsCreated:    in.Metadata.DCTermsCreated,
			DCDateCreated:     in.Metadata.DCDateCreated,
			DCDate:            in.Metadata.DCDate,
			DCTermsType:       in.Metadata.DCTermsType,
			DCType:            in.Metadata.DCType,
			DCTermsAudience:   in.Metadata.DCTermsAudience,
			DCTermsSubject:    in.Metadata.DCTermsSubject,
			DCSubject:         in.Metadata.DCSubject,
			DCDescription:     in.Metadata.Description,
			DCTermsKeywords:   in.Metadata.DCTermsKeywords,
			ModifiedTime:      in.Metadata.ModifiedTime,
			PublishedTime:     in.Metadata.PublishedTime,
			ArticleTag:        in.Metadata.ArticleTag,
			ArticleSection:    in.Metadata.ArticleSection,
			SourceURL:         in.Metadata.SourceURL,
			StatusCode:        in.Metadata.StatusCode,
			Error:             in.Metadata.Error,
			ExtraFields:       in.Metadata.ExtraFields,
		},
	}
}
func (c *Crawl) BatchScrapeUrls(ctx context.Context, urls []string, params ScrapeParams) (scrapeResponse *ScrapeResponse, err error) {
	response, err := crawl.ClientInterface.BatchScrapeUrls(ctx, &models.ScrapeOptionsMultiple{
		Url:             urls,
		Formats:         params.Formats,
		Headers:         params.Headers,
		IncludeTags:     params.IncludeTags,
		ExcludeTags:     params.ExcludeTags,
		OnlyMainContent: params.OnlyMainContent,
		WaitFor:         params.WaitFor,
		Timeout:         params.Timeout,
		BrowserOptions: models.ICreateBrowser{
			SessionName: params.BrowserOptions.SessionName,
		},
	})
	if err != nil {
		return nil, err
	}
	return &ScrapeResponse{
		ID:          response.ID,
		InvalidURLs: response.InvalidURLs,
	}, nil
}

func (c *Crawl) CheckBatchScrapeStatus(ctx context.Context, id string) (scrapeStatusResponseMultiple *ScrapeStatusResponseMultiple, err error) {
	response, err := crawl.ClientInterface.CheckBatchScrapeStatus(ctx, id)
	if err != nil {
		return nil, err
	}
	scrapeStatusResponseMultiple = c.internalScrapeStatusResponseMultipleFormat(response)
	return
}

func (c *Crawl) internalScrapeStatusResponseMultipleFormat(in *models.ScrapeStatusResponseMultiple) *ScrapeStatusResponseMultiple {
	multiple := &ScrapeStatusResponseMultiple{
		Completed: in.Completed,
		Total:     in.Total,
		Status:    Status(in.Status),
	}
	for _, datum := range in.Data {
		multiple.Data = append(multiple.Data, *c.internalScrapingCrawlDocumentFormat(&datum))
	}
	return multiple
}

func (c *Crawl) AsyncCrawlUrl(ctx context.Context, url string, params CrawlParams) (id string, err error) {
	crawlUrl, err := crawl.ClientInterface.CrawlUrl(ctx, &models.CrawlParams{
		Url:                    url,
		IncludePaths:           params.IncludePaths,
		ExcludePaths:           params.ExcludePaths,
		MaxDepth:               params.MaxDepth,
		MaxDiscoveryDepth:      params.MaxDiscoveryDepth,
		Limit:                  params.Limit,
		AllowBackwardLinks:     params.AllowBackwardLinks,
		AllowExternalLinks:     params.AllowExternalLinks,
		IgnoreSitemap:          params.IgnoreSitemap,
		DeduplicateSimilarURLs: params.DeduplicateSimilarURLs,
		IgnoreQueryParameters:  params.IgnoreQueryParameters,
		RegexOnFullURL:         params.RegexOnFullURL,
		Delay:                  params.Delay,
		ScrapeOptions: models.CrawlScrapeOptions{
			Formats:         params.ScrapeOptions.Formats,
			Headers:         params.ScrapeOptions.Headers,
			IncludeTags:     params.ScrapeOptions.IncludeTags,
			ExcludeTags:     params.ScrapeOptions.ExcludeTags,
			OnlyMainContent: params.ScrapeOptions.OnlyMainContent,
			WaitFor:         params.ScrapeOptions.WaitFor,
			Timeout:         params.ScrapeOptions.Timeout,
		},
		BrowserOptions: models.ICreateBrowser{
			SessionName:      params.BrowserOptions.SessionName,
			SessionTTL:       params.BrowserOptions.SessionTTL,
			SessionRecording: params.BrowserOptions.SessionRecording,
			ProxyCountry:     params.BrowserOptions.ProxyCountry,
			ProxyURL:         params.BrowserOptions.ProxyURL,
			Fingerprint:      params.BrowserOptions.Fingerprint,
		},
	})
	if err != nil {
		return "", err
	}
	return crawlUrl, nil
}

func (c *Crawl) CrawlUrl(ctx context.Context, url string, params CrawlParams) (crawlStatusResponse *CrawlStatusResponse, err error) {
	id, err := c.AsyncCrawlUrl(ctx, url, params)
	if err != nil {
		return nil, err
	}
	for {
		crawlStatusResponse, err = c.CheckCrawlStatus(ctx, id)
		if err != nil {
			return nil, err
		}
		switch crawlStatusResponse.Status {
		case StatusCompleted:
			return crawlStatusResponse, nil
		case StatusActive, StatusPaused, StatusPending, StatusQueued, StatusWaiting, StatusScraping:
			log.Info("Scraping status: ", crawlStatusResponse.Status)
			time.Sleep(time.Millisecond * 400)
		default:
			return nil, errors.New(fmt.Sprintf("Crawl job failed or was stopped. Status: %s", crawlStatusResponse.Status))
		}
	}

}
func (c *Crawl) CheckCrawlStatus(ctx context.Context, id string) (crawlStatusResponse *CrawlStatusResponse, err error) {
	response, err := crawl.ClientInterface.CheckCrawlStatus(ctx, id)
	if err != nil {
		return nil, err
	}
	var scrapingCrawlDocuments []ScrapingCrawlDocument
	for _, datum := range response.Data {
		format := c.internalScrapingCrawlDocumentFormat(&datum)
		scrapingCrawlDocuments = append(scrapingCrawlDocuments, *format)
	}
	return &CrawlStatusResponse{
		Status: Status(response.Status),
		Data:   scrapingCrawlDocuments,
	}, nil
}
func (c *Crawl) CheckCrawlErrors(ctx context.Context, id string) (crawlErrorsResponse *CrawlErrorsResponse, err error) {
	response, err := crawl.ClientInterface.CheckCrawlErrors(ctx, id)
	if err != nil {
		return nil, err
	}
	var robotsBlocked []string
	var crawlErrorDetail []CrawlErrorDetail
	for _, detail := range response.Errors {
		crawlErrorDetail = append(crawlErrorDetail, CrawlErrorDetail{
			ID:        detail.ID,
			Timestamp: detail.Timestamp,
			Url:       detail.Url,
			Error:     detail.Error,
		})
	}
	for _, rb := range response.RobotsBlocked {
		robotsBlocked = append(robotsBlocked, rb)
	}
	return &CrawlErrorsResponse{
		Errors:        crawlErrorDetail,
		RobotsBlocked: robotsBlocked,
	}, nil
}
func (c *Crawl) CancelCrawl(ctx context.Context, id string) (success bool, err error) {
	_, err = crawl.ClientInterface.CancelCrawl(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *Crawl) Close() error {
	return nil
}
