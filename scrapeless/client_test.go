package scrapeless

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/proxies"
	"testing"
)

func TestNew(t *testing.T) {
	client := New(WithProxy())
	proxy, err := client.Proxy.Proxy(context.Background(), proxies.ProxyActor{
		Country:         "",
		SessionDuration: 0,
		SessionId:       "",
		Gateway:         "",
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(proxy)
	//namespaces, err := client.Storage.KV.ListNamespaces(context.Background(), 1, 10, true)
	//if err != nil {
	//	t.Error(err.Error())
	//	return
	//}
	//t.Logf("%v", namespaces)

	//once, err := client.Browser.CreateOnce(context.Background(), browser.ActorOnce{
	//	Input:        browser.Input{},
	//	ProxyCountry: "",
	//})
	//if err != nil {
	//	t.Error(err.Error())
	//	return
	//}
	//t.Logf("%v", once)

	//captchaTaskId, err := client.Captcha.Create(context.TODO(), &captcha.CaptchaSolverReq{
	//	Actor: "captcha.recaptcha",
	//	Input: captcha.Input{
	//		Version: captcha.RecaptchaVersionV2,
	//		PageURL: "https://venue.cityline.com",
	//		SiteKey: "6Le_J04UAAAAAIAfpxnuKMbLjH7ISXlMUzlIYwVw",
	//	},
	//	Proxy: captcha.ProxyInfo{
	//		Country: "US",
	//	},
	//})
	//if err != nil {
	//	log.Error(err.Error())
	//}
	//log.Infof("%v", captchaTaskId)
	//// Wait for captcha task to be solved
	//time.Sleep(time.Second * 20)
	//captchaResult, err := client.Captcha.ResultGet(context.TODO(), &captcha.CaptchaSolverReq{
	//	TaskId: captchaTaskId,
	//})
	//if err != nil {
	//	log.Error(err.Error())
	//}
	//log.Infof("%v", captchaResult)
}
