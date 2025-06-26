package crawl

import (
	"context"
	crawl_http "github.com/smash-hq/sdk-go/internal/remote/crawl/http"
	"github.com/smash-hq/sdk-go/internal/remote/crawl/models"
	"github.com/smash-hq/sdk-go/scrapeless/log"
)

type Captcha interface {
	ScrapeUrl(ctx context.Context, req *models.ScrapeOptions) (id string, err error)
	BatchScrapeUrls(ctx context.Context, req *models.ScrapeOptionsMultiple) (scrapeResponse *models.ScrapeResponse, err error)
	CheckScrapeStatus(ctx context.Context, id string) (scrapeStatusResponse *models.ScrapeStatusResponse, err error)
	CheckBatchScrapeStatus(ctx context.Context, id string) (scrapeStatusResponseMultiple *models.ScrapeStatusResponseMultiple, err error)
	CrawlUrl(ctx context.Context, req *models.CrawlParams) (id string, err error)
	CheckCrawlStatus(ctx context.Context, id string) (crawlStatusResponse *models.CrawlStatusResponse, err error)
	CheckCrawlErrors(ctx context.Context, id string) (crawlErrorsResponse *models.CrawlErrorsResponse, err error)
	CancelCrawl(ctx context.Context, id string) (errorResponse *models.ErrorResponse, err error)
}

var ClientInterface Captcha

func NewClient(serverMode, baseUrl string) {
	switch serverMode {
	case "grpc":
		log.Info("grpc...")
	case "dev":
		log.Info("dev...")
	default:
		crawl_http.Init(baseUrl)
		ClientInterface = crawl_http.Default()
	}
}
