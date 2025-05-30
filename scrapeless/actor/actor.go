package actor

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/browser"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/captcha"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/httpserver"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/proxies"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/router"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/storage"
	"github.com/spf13/viper"
	"reflect"
)

type Actor struct {
	Browser  browser.Browser
	Proxy    proxies.Proxy
	Captcha  captcha.Captcha
	Storage  storage.Storage
	server   httpserver.Server
	Router   router.Router
	closeFun []func() error
}

// New creates a new Actor.
func New() *Actor {
	var actor = new(Actor)
	actor.Storage = storage.NewStorageHttp()
	actor.Browser = browser.NewBHttp()
	actor.Captcha = captcha.NewCaHttp()
	actor.Proxy = proxies.NewPHttp()
	actor.server = httpserver.New()
	actor.Router = router.New()
	return actor
}

// Close closes the actor.
func (a *Actor) Close() {
	for _, f := range a.closeFun {
		_ = f()
	}
}

// Input get input data from env.
func (a *Actor) Input(data any) error {
	input := viper.GetStringMapString("SCRAPELESS_INPUT")
	inputData, _ := json.Marshal(input)
	tf := reflect.TypeOf(data)
	if tf.Kind() != reflect.Ptr {
		return errors.New("data must be ptr")
	}
	return json.Unmarshal(inputData, data)
}

func (a *Actor) Start() error {
	return a.server.Start(fmt.Sprintf(":%s", env.Env.Actor.HttpPort))
}
