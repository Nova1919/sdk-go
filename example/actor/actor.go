package main

import (
	"context"
	"encoding/json"
	"github.com/scrapeless-ai/sdk-go/scrapeless/actor"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/browser"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/captcha"
	proxy2 "github.com/scrapeless-ai/sdk-go/scrapeless/services/proxies"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/storage/queue"
	"net/url"
)

type Input struct {
	//...
}

/**
 * Complete Actor Usage Example
 * Demonstrates how to use various features of the Actor struct
 */
func main() {
	// 1. Initialize Actor
	a := actor.New()
	defer a.Close()

	// 2. Get input data
	input := Input{}
	err := a.Input(&input)
	if err != nil {
		panic(err)
	}
	log.Info("input:", input)

	// 3. Dataset operations
	// 3.1 Add items to dataset
	ok, err := a.Storage.GetDataset().AddItems(context.Background(), []map[string]any{{"title": "test"}})
	if err != nil {
		log.Error(err)
	}
	log.Info("ok:", ok)
	// 3.2  Get data from dataset
	items, err := a.Storage.GetDataset().GetItems(context.Background(), 1, 1, false)
	if err != nil {
		log.Error(err.Error())
	}
	log.Infof("%v", items)
	// 3.3 You can also use the underlying API directly
	id, name, err := a.Storage.GetDataset().CreateDataset(context.Background(), "New Dataset")
	if err != nil {
		log.Error(err.Error())
	}
	log.Infof("id:%v, name:%v", id, name)

	// 4. KV storage operations
	// 4.1 Store data
	value, err := a.Storage.GetKv().SetValue(context.Background(), "key", "value", 0)
	if err != nil {
		log.Error(err.Error())
	}
	log.Infof("value:%v", value)
	// 4.2 Retrieve data
	kvValue, err := a.Storage.GetKv().GetValue(context.Background(), "key")
	if err != nil {
		log.Error(err.Error())
	}
	log.Infof("kvValue:%v", kvValue)
	// 4.3  Store structured data
	m := map[string]any{"name": "test"}
	kvData, _ := json.Marshal(m)
	setValue, err := a.Storage.GetKv().SetValue(context.Background(), "key", string(kvData), 0)
	if err != nil {
		log.Error(err.Error())
	}
	log.Infof("setValue:%v", setValue)
	// 4.4 Retrieve and parse structured data
	kvValue, err = a.Storage.GetKv().GetValue(context.Background(), "key")
	if err != nil {
		log.Error(err.Error())
	}
	structData := map[string]any{}
	if err = json.Unmarshal([]byte(kvValue), &structData); err != nil {
		log.Error(err.Error())
	}
	log.Infof("structData:%v", structData)

	// 5. Object storage operations
	// In a real environment, you can upload files here
	// Due to example environment limitations, only showing the code
	//// 5.1 Upload file to object storage
	//objId, err := a.Storage.GetObject().Put(context.Background(), "object.json", []byte("data"))
	//if err != nil {
	//	log.Error(err.Error())
	//}
	//log.Infof("objId:%v", objId)
	//// 5.2 Retrieve file from object storage
	//get, err := a.Storage.GetObject().Get(context.Background(), objId)
	//if err != nil {
	//	log.Error(err.Error())
	//}
	//log.Infof("get:%v", string(get))

	// 6. Queue operations
	// 6.1 Push message to queue
	msgId, err := a.Storage.GetQueue().Push(context.Background(), queue.PushQueue{
		Name:    "example",
		Payload: []byte("example"),
	})
	if err != nil {
		log.Error(err.Error())
	}
	log.Infof("msgId:%v", msgId)
	// 6.2 Pull message from queue
	queueResp, err := a.Storage.GetQueue().Pull(context.Background(), 1)
	if err != nil {
		log.Error(err.Error())
	}
	log.Infof("%v", queueResp)
	// 6.3 Acknowledge message completion
	err = a.Storage.GetQueue().Ack(context.Background(), queueResp[0].QueueID)
	if err != nil {
		log.Error(err.Error())
	}
	// 7. CAPTCHA solving
	solver, err := a.Captcha.Solver(context.Background(), &captcha.CaptchaSolverReq{
		Actor: "captcha.recaptcha",
		Input: captcha.Input{
			Version: captcha.RecaptchaVersionV2,
			PageURL: "https://venue.cityline.com",
			SiteKey: "6Le_J04UAAAAAIAfpxnuKMbLjH7ISXlMUzlIYwVw",
		},
		Proxy: captcha.ProxyInfo{
			Country: "US",
		},
	})
	if err != nil {
		log.Error(err.Error())
	}
	log.Infof("%v", solver)

	// 8. Proxy usage
	proxy, err := a.Proxy.Proxy(context.Background(), proxy2.ProxyActor{
		Country:         "US",
		SessionDuration: 180,
		SessionId:       "YOUR SESSION ID",
		Gateway:         "YOU GATEWAY",
	})
	if err != nil {
		log.Error(err.Error())
	}
	log.Info(proxy)
	parse, err := url.Parse(proxy)
	if err != nil {
		log.Error(err.Error())
	}
	// set proxy
	// client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(parse)}}
	log.Info("proxy:", parse.String())

	// 9. Browser
	browserInfo, err := a.Browser.Create(context.Background(), browser.Actor{
		Input:        browser.Input{SessionTtl: "180"},
		ProxyCountry: "US",
	})
	if err != nil {
		log.Error(err.Error())
	}
	log.Infof("%+v", browserInfo)
}
