# Scrapeless SDK Go

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**[English](README.md) | [中文文档](README-zh.md)**

[Scrapeless AI](https://scrapeless.com) 官方 Go 语言 SDK - 一个强大的网页抓取和浏览器自动化平台，帮助您大规模提取任何网站的数据。

## 📑 目录

- [🌟 特性](#-特性)
- [📦 安装](#-安装)
- [🚀 快速开始](#-快速开始)
- [📖 使用示例](#-使用示例)
- [🔧 API 参考](#-api-参考)
- [📚 示例](#-示例)
- [🛠️ 贡献&开发指南](#-贡献开发指南)
- [📄 许可证](#-许可证)
- [📞 支持](#-支持)
- [🏢 关于 Scrapeless](#-关于-scrapeless)

## 🌟 特性

- **浏览器自动化**：支持远程浏览器会话
- **网页抓取**：通过智能解析从任何网站提取数据
- **SERP 抓取**：高精度提取搜索引擎结果
- **代理管理**：内置代理轮换和地理定位
- **Actor 系统**：在云端运行自定义自动化脚本
- **存储解决方案**：为您的抓取项目提供持久化数据存储

## 📦 安装

使用 `go get` 安装 SDK：

```bash
go get -u github.com/scrapeless-ai/sdk-go
```


## 🚀 快速开始

### 基本设置

```
package main

import (
	scrapeless "github.com/scrapeless-ai/sdk-go/scrapeless/actor"
)

func main() {
  	// 初始化 actor
	actor := scrapeless.New()
	defer actor.Close()
}
```

### 环境变量

您也可以使用环境变量配置 SDK：

```bash
# 必需
SCRAPELESS_API_KEY=your-api-key

# 可选 - 自定义 API 端点
SCRAPELESS_BASE_API_URL=https://api.scrapeless.com
SCRAPELESS_ACTOR_API_URL=https://actor.scrapeless.com
SCRAPELESS_STORAGE_API_URL=https://storage.scrapeless.com
SCRAPELESS_BROWSER_API_URL=https://browser.scrapeless.com
SCRAPELESS_CRAWL_API_URL=https://crawl.scrapeless.com
```

## 📖 使用示例

### 浏览器自动化

```go
package main

import (
	"context"
	scrapeless "github.com/scrapeless-ai/sdk-go/scrapeless/actor"
	"github.com/scrapeless-ai/sdk-go/scrapeless/browser"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
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

### 网页抓取

```go
package main

import (
	"context"
	scrapeless "github.com/scrapeless-ai/sdk-go/scrapeless/actor"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"github.com/scrapeless-ai/sdk-go/scrapeless/scraping"
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

### SERP 抓取

```go
package main

import (
	"context"
	scrapeless "github.com/scrapeless-ai/sdk-go/scrapeless/actor"
	"github.com/scrapeless-ai/sdk-go/scrapeless/deepserp"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
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

### Actor 系统

```go
package main

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/internal/remote/actor"
	"github.com/scrapeless-ai/sdk-go/scrapeless"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
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

## 🔧 API 参考

### 可用服务

SDK 提供以下服务：

- `Client.Browser` - 浏览器会话管理
- `Client.Scraping` - 网页抓取和数据提取
- `Client.DeepSerp` - 搜索引擎结果提取
- `Client.Universal` - 通用数据提取
- `Client.Proxy` - 代理管理
- `Client.Actor` - 自定义自动化的 Actor 系统
- `Client.Storage` - 数据存储解决方案
- `Client.Server` - http服务
- `Client.Router` - 路由访问
- `Client.Captcha` - 验证码处理


## 📚 示例

查看 `example` 目录获取完整的使用示例：

- [Actor 系统](./example/actor_service/actor_service.go)
- [SERP 抓取](./example/deepserp/deepserp.go)
- [网页抓取](./example/scraping/scraping.go)
- [浏览器操作示例](./example/browser/browser.go)
- [验证码识别示例](./example/captcha/captcha.go)
- [代理管理示例](./example/proxy/proxy.go)
- [存储dataset使用示例](./example/storage_dataset/storage_dataset.go)
- [存储kv使用示例](./example/storage_kv/storage_kv.go)
- [存储object使用示例](./example/storage_object/storage_object.go)
- [存储queue使用示例](./example/storage_queue/storage_queue.go)
- [路由调用](./example/router/router.go)
- [http服务](./example/httpserver/httpserver.go)

## 🛠️ 贡献&开发指南

欢迎所有形式的贡献！关于如何提交 issue、PR、代码规范、本地开发等详细内容，请参见[贡献与开发指南](./CONTRIBUTING-zh.md)。

**快速开始：**

```bash
git clone https://github.com/scrapeless-ai/sdk-go.git
cd scrapeless-actor-sdk-go
go mod tidy
go run ./example/actor/actor.go
```

更多项目结构、最佳实践等内容请参见 [CONTRIBUTING-zh.md](./CONTRIBUTING-zh.md)。

## 📄 许可证

本项目基于 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 📞 支持

- 📖 **文档**: [https://docs.scrapeless.com](https://docs.scrapeless.com)
- 💬 **社区**: [加入我们的 Discord](https://backend.scrapeless.com/app/api/v1/public/links/discord)
- 🐛 **问题**: [GitHub Issues](https://github.com/scrapeless-ai/scrapeless-sdk-node/issues)
- 📧 **邮箱**: [support@scrapeless.com](mailto:support@scrapeless.com)


## 🏢 关于 Scrapeless

Scrapeless 是一个强大的网页抓取和浏览器自动化平台，帮助企业大规模从任何网站提取数据。我们的平台提供：

- 高性能网页抓取基础设施
- 全球代理网络
- 浏览器自动化功能
- 企业级可靠性和支持

访问 [scrapeless.com](https://scrapeless.com) 了解更多并开始使用。

---

由 Scrapeless 团队用 ❤️ 制作


        