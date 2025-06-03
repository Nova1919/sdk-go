package scrapeless

import (
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/actor"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/browser"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/captcha"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/deepserp"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/httpserver"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/proxies"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/router"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/scraping"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/storage"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/universal"
)

type Client struct {
	Browser   browser.Browser
	Proxy     proxies.Proxy
	Captcha   captcha.Captcha
	Storage   storage.Storage
	Server    httpserver.Server
	Router    router.Router
	DeepSerp  deepserp.Deepserp
	Scraping  scraping.Scraping
	Universal universal.Universal
	Actor     actor.ActorService
	CloseFun  []func() error
}

func New(opts ...Option) *Client {
	var client = new(Client)
	for _, opt := range opts {
		opt.Apply(client)
	}
	client.Router = router.New()
	return client
}

// Close closes the Client.
func (c *Client) Close() {
	for _, f := range c.CloseFun {
		_ = f()
	}
}

const (
	typeHttp = "http"
	typeGrpc = "grpc"
)

type Option interface {
	Apply(*Client)
}

type BrowserOption struct {
	tp string
}

func (o *BrowserOption) Apply(c *Client) {
	if o.tp == typeGrpc {
		c.Browser = browser.NewBGrpc()
		c.CloseFun = append(c.CloseFun, c.Browser.Close)
	} else {
		c.Browser = browser.NewBHttp()
		c.CloseFun = append(c.CloseFun, c.Browser.Close)
	}
}

// WithBrowser choose browser type.
func WithBrowser(tp ...string) Option {
	if len(tp) == 0 {
		tp = append(tp, typeHttp)
	}
	return &BrowserOption{
		tp: tp[0],
	}
}

type ProxyOption struct {
	tp string
}

func (o *ProxyOption) Apply(a *Client) {
	if o.tp == typeGrpc {
		a.CloseFun = append(a.CloseFun, a.Proxy.Close)
	} else {
		a.Proxy = proxies.NewPHttp()
		a.CloseFun = append(a.CloseFun, a.Proxy.Close)
	}
}

// WithProxy choose proxies type.
func WithProxy(tp ...string) Option {
	if len(tp) == 0 {
		tp = append(tp, typeHttp)
	}
	return &ProxyOption{
		tp: tp[0],
	}
}

type CaptchaOption struct {
	tp string
}

func (o *CaptchaOption) Apply(a *Client) {
	if o.tp == typeGrpc {
		a.CloseFun = append(a.CloseFun, a.Captcha.Close)
	} else {
		a.Captcha = captcha.NewCaHttp()
		a.CloseFun = append(a.CloseFun, a.Captcha.Close)
	}
}

// WithCaptcha choose captcha type.
func WithCaptcha(tp ...string) Option {
	if len(tp) == 0 {
		tp = append(tp, typeHttp)
	}
	return &CaptchaOption{
		tp: tp[0],
	}
}

type StorageOption struct {
	tp string
}

func (o *StorageOption) Apply(a *Client) {
	if o.tp == typeGrpc {
		a.CloseFun = append(a.CloseFun, a.Storage.Close)
	} else {
		a.Storage = storage.NewStorage()
		a.CloseFun = append(a.CloseFun, a.Storage.Close)
	}
}

// WithStorage choose storage type.
func WithStorage(tp ...string) Option {
	if len(tp) == 0 {
		tp = append(tp, typeHttp)
	}
	return &StorageOption{
		tp: tp[0],
	}
}

type ServerOption struct {
	mode httpserver.ServerMode
}

func (s *ServerOption) Apply(a *Client) {
	a.Server = httpserver.New(s.mode)
}

// WithServer choose server mode.
func WithServer(mode ...httpserver.ServerMode) Option {
	if len(mode) == 0 {
		mode = append(mode, httpserver.ReleaseMode)
	}
	return &ServerOption{mode: mode[0]}
}

type DeepSerpOption struct {
	tp string
}

func (d *DeepSerpOption) Apply(c *Client) {
	if d.tp == typeGrpc {
		c.CloseFun = append(c.CloseFun, c.DeepSerp.Close)
	} else {
		c.DeepSerp = deepserp.New()
		c.CloseFun = append(c.CloseFun, c.DeepSerp.Close)
	}
}

// WithDeepSerp choose DeepSerp type.
func WithDeepSerp(tp ...string) Option {
	if len(tp) == 0 {
		tp = append(tp, typeHttp)
	}
	return &DeepSerpOption{
		tp: tp[0],
	}
}

type ScrapingOption struct {
	tp string
}

func (s *ScrapingOption) Apply(c *Client) {
	if s.tp == typeGrpc {
		c.CloseFun = append(c.CloseFun, c.Scraping.Close)
	} else {
		c.Scraping = scraping.New()
		c.CloseFun = append(c.CloseFun, c.Scraping.Close)
	}
}

// WithScraping choose scraping type.
func WithScraping(tp ...string) Option {
	if len(tp) == 0 {
		tp = append(tp, typeHttp)
	}
	return &ScrapingOption{
		tp: tp[0],
	}
}

type UniversalOption struct {
	tp string
}

func (s *UniversalOption) Apply(c *Client) {
	if s.tp == typeGrpc {
		c.CloseFun = append(c.CloseFun, c.Universal.Close)
	} else {
		c.Universal = universal.New()
		c.CloseFun = append(c.CloseFun, c.Universal.Close)
	}
}

// WithUniversal choose universal type.
func WithUniversal(tp ...string) Option {
	if len(tp) == 0 {
		tp = append(tp, typeHttp)
	}
	return &UniversalOption{
		tp: tp[0],
	}
}

type ActorOption struct {
	tp string
}

func (s *ActorOption) Apply(c *Client) {
	if s.tp == typeGrpc {
		c.CloseFun = append(c.CloseFun, c.Actor.Close)
	} else {
		c.Actor = actor.NewActorHttp()
		c.CloseFun = append(c.CloseFun, c.Actor.Close)
	}
}

// WithActor choose Actor type.
func WithActor(tp ...string) Option {
	if len(tp) == 0 {
		tp = append(tp, typeHttp)
	}
	return &ActorOption{
		tp: tp[0],
	}
}
