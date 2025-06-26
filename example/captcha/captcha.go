package main

import (
	"context"
	"github.com/smash-hq/sdk-go/scrapeless"
	"github.com/smash-hq/sdk-go/scrapeless/log"
	"github.com/smash-hq/sdk-go/scrapeless/services/captcha"
	"time"
)

func main() {
	client := scrapeless.New(scrapeless.WithCaptcha())
	//Create captcha task
	captchaTaskId, err := client.Captcha.Create(context.TODO(), &captcha.CaptchaSolverReq{
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
	log.Infof("%v", captchaTaskId)
	// Wait for captcha task to be solved
	time.Sleep(time.Second * 20)
	captchaResult, err := client.Captcha.ResultGet(context.TODO(), &captcha.CaptchaSolverReq{
		TaskId: captchaTaskId,
	})
	if err != nil {
		log.Error(err.Error())
	}
	log.Infof("%v", captchaResult)
}
