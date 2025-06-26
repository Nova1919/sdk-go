# Scrapeless SDK Go

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)



The official Go SDK of [Scrapeless AI](https://scrapeless.com) - a powerful web scraping and browser automation platform that helps you extract data from any website at scale.

## 📑 Table of Contents

- [🌟 Features](#-features)
- [📦 Installation](#-installation)
- [🚀 Quick Start](#-quick-start)
- [📖 Usage Examples](#-usage-examples)
- [🔧 API Reference](#-api-reference)
- [📚 Examples](#-examples)
- [🛠️ Contribution & Development Guide](#-contribution--development-guide)
- [📄 License](#-license)
- [📞 Support](#-support)
- [🏢 About Scrapeless](#-about-scrapeless)

## 🌟 Features

- **Browser Automation**: Supports remote browser sessions.
- **Web Scraping**: Extracts data from any website through intelligent parsing.
- **SERP Scraping**: Extracts search engine results with high accuracy.
- **Proxy Management**: Built-in proxy rotation and geolocation.
- **Actor System**: Runs custom automation scripts in the cloud.
- **Storage Solutions**: Provides persistent data storage for your scraping projects.

## 📦 Installation

Install the SDK using `go get`:

```bash
go get -u github.com/smash-hq/sdk-go
```

## 🚀 Quick Start

### Basic Setup

```go
package main

import (
	scrapeless "github.com/smash-hq/sdk-go/scrapeless/actor"
)

func main() {
	// Initialize the actor
	actor := scrapeless.New()
	defer actor.Close()
}
```

### Environment Variables

You can also configure the SDK using environment variables:

```bash
# Required
SCRAPELESS_API_KEY=your-api-key

# Optional - Custom API endpoints
SCRAPELESS_BASE_API_URL=https://api.scrapeless.com
SCRAPELESS_ACTOR_API_URL=https://actor.scrapeless.com
SCRAPELESS_STORAGE_API_URL=https://storage.scrapeless.com
SCRAPELESS_BROWSER_API_URL=https://browser.scrapeless.com
SCRAPELESS_CRAWL_API_URL=https://crawl.scrapeless.com
```

## 📖 Usage Examples

### Browser Automation

```go
package main

import (
	"context"
	scrapeless "github.com/smash-hq/sdk-go/scrapeless/actor"
	"github.com/smash-hq/sdk-go/scrapeless/browser"
	"github.com/smash-hq/sdk-go/scrapeless/log"
)

func main() {
	client := scrapeless.New(scrapeless.WithBrowser())
	defer client.Close()

	browserInfo, err := client.Browser.Create(context.Background(), browser.Actor{
		Input:        browser.Input{SessionTtl: "180"},
		ProxyCountry: "US",
	})
	if err != nil {
		panic(err)
	}
	log.Infof("%+v", browserInfo)
}
```

### Web Scraping

```go
package main

import (
	"context"
	scrapeless "github.com/smash-hq/sdk-go/scrapeless/actor"
	"github.com/smash-hq/sdk-go/scrapeless/log"
	"github.com/smash-hq/sdk-go/scrapeless/scraping"
)

func main() {
	client := scrapeless.New(scrapeless.WithScraping())

	scrape, err := client.Scraping.Scrape(context.Background(), scraping.ScrapingTaskRequest{
		Actor: "scraper.google.search",
		Input: map[string]interface{}{
			"q": "nike site:www.nike.com",
		},
		ProxyCountry: "US",
	})
	if err != nil {
		log.Errorf("scraping create err:%v", err)
		return
	}
	log.Infof("%+v", scrape)
}
```

### SERP Scraping

```go
package main

import (
	"context"
	scrapeless "github.com/smash-hq/sdk-go/scrapeless/actor"
	"github.com/smash-hq/sdk-go/scrapeless/deepserp"
	"github.com/smash-hq/sdk-go/scrapeless/log"
)

func main() {
	client := scrapeless.New(scrapeless.WithDeepSerp())

	scrape, err := client.DeepSerp.Scrape(context.Background(), deepserp.DeepserpTaskRequest{
		Actor: "scraper.google.search",
		Input: map[string]interface{}{
			"q": "nike site:www.nike.com",
		},
		ProxyCountry: "US",
	})
	if err != nil {
		log.Errorf("scraping create err:%v", err)
		return
	}
	log.Infof("%+v", scrape)
}
```

### Actor System

```go
package main

import (
	"context"
	"github.com/smash-hq/sdk-go/internal/remote/actor"
	"github.com/smash-hq/sdk-go/scrapeless"
	"github.com/smash-hq/sdk-go/scrapeless/log"
)

func main() {
	client := scrapeless.New(scrapeless.WithActor())
	defer client.Close()

	runId, err := client.Actor.Run(context.Background(), actor.IRunActorData{
		ActorId: "554bbd68-c787-4900-b8b2-1086369c96e1",
		Input: map[string]string{
			"name": "test",
			"url":  "https://www.google.com",
		},
		RunOptions: actor.RunOptions{
			Version: "v0.0.3",
		},
	})
	if err != nil {
		panic(err)
	}
	runInfo, err := client.Actor.GetRunInfo(context.Background(), runId)
	if err != nil {
		panic(err)
	}
	log.Infof("runInfo:%+v", runInfo)
}
```
### Crawl

```go
package main

import (
	"context"
	"github.com/smash-hq/sdk-go/scrapeless"
	"github.com/smash-hq/sdk-go/scrapeless/log"
	"github.com/smash-hq/sdk-go/scrapeless/services/crawl"
)

func main() {
	client := scrapeless.New(scrapeless.WithCrawl())

	// Crawl
	response, err := client.Crawl.CrawlUrl(context.Background(), "https://redditinc.com/blog", crawl.CrawlParams{
		Limit: 10,
		ScrapeOptions: crawl.ScrapeOptions{
			Formats: []string{"links",
				"markdown",
				"html",
				"screenshot"},
		},
		BrowserOptions: crawl.ICreateBrowser{
			SessionName:      "Crawl",
			SessionTTL:       "900",
			SessionRecording: "true",
			ProxyCountry:     "ANY",
		},
	})
	if err != nil {
		panic(err)
	}
	log.Infof("Crawl response: %v", response)

	// scrape
	scrapeResponse, err := client.Crawl.ScrapeUrl(context.Background(), "https://docs.scrapeless.com/en/overview/", crawl.ScrapeOptions{
		BrowserOptions: crawl.ICreateBrowser{
			SessionName:      "Crawl",
			SessionTTL:       "900",
			SessionRecording: "true",
			ProxyCountry:     "ANY",
		},
	})
	if err != nil {
		panic(err)
	}
	log.Infof("Scrape response: %v", scrapeResponse)
}
```

## 🔧 API Reference

### Available Services

The SDK provides the following services:

- `Client.Browser` - Browser session management.
- `Client.Scraping` - Web scraping and data extraction.
- `Client.DeepSerp` - Search engine result extraction.
- `Client.Universal` - Universal data extraction.
- `Client.Proxy` - Proxy management.
- `Client.Actor` - Actor system for custom automation.
- `Client.Storage` - Data storage solutions.
- `Client.Server` - HTTP service.
- `Client.Router` - Route access.
- `Client.Captcha` - Captcha processing.

## 📚 Examples

Check the `example` directory for complete usage examples:

- [Actor System](./example/actor_service/actor_service.go)
- [SERP Scraping](./example/deepserp/deepserp.go)
- [Web Scraping](./example/scraping/scraping.go)
- [Browser Operation Example](./example/browser/browser.go)
- [Captcha Recognition Example](./example/captcha/captcha.go)
- [Proxy Management Example](./example/proxy/proxy.go)
- [Storage Dataset Usage Example](./example/storage_dataset/storage_dataset.go)
- [Storage KV Usage Example](./example/storage_kv/storage_kv.go)
- [Storage Object Usage Example](./example/storage_object/storage_object.go)
- [Storage Queue Usage Example](./example/storage_queue/storage_queue.go)
- [Route Call](./example/router/router.go)
- [HTTP Service](./example/httpserver/httpserver.go)
- [Crawl Usage Example](./example/crawl/crawl.go)

## 🛠️ Contribution & Development Guide

All forms of contributions are welcome! For detailed information on how to submit issues, PRs, code specifications, local development, etc., please refer to the [Contribution & Development Guide](./CONTRIBUTING.md).

**Quick Start**:

```bash
git clone https://github.com/smash-hq/sdk-go.git
cd scrapeless-actor-sdk-go
go mod tidy
go run ./example/actor/actor.go
```

For more information on project structure, best practices, etc., please refer to [CONTRIBUTING.md](./CONTRIBUTING.md).

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Support

- 📖 **Documentation**: [https://docs.scrapeless.com](https://docs.scrapeless.com)
- 💬 **Community**: [Join our Discord](https://backend.scrapeless.com/app/api/v1/public/links/discord)
- 🐛 **Issues**: [GitHub Issues](https://github.com/scrapeless-ai/scrapeless-sdk-node/issues)
- 📧 **Email**: [support@scrapeless.com](mailto:support@scrapeless.com)

## 🏢 About Scrapeless

Scrapeless is a powerful web scraping and browser automation platform that helps enterprises extract data from any website at scale. Our platform provides:

- High-performance web scraping infrastructure.
- Global proxy network.
- Browser automation capabilities.
- Enterprise-level reliability and support.

Visit [scrapeless.com](https://scrapeless.com) to learn more and get started.

---

Made with ❤️ by the Scrapeless team

        