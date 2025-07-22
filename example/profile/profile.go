package main

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/scrapeless"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/browser"
)

func main() {
	client := scrapeless.New(scrapeless.WithProfile(), scrapeless.WithBrowser())

	// Create profile
	response, err := client.Profile.CreateProfile(context.Background(), "my_profile")
	if err != nil {
		panic(err)
	}
	log.Infof("create profile response: %v", response)

	browserInfo, err := client.Browser.Create(context.Background(), browser.Actor{
		//Input:          browser.Input{SessionTtl: "180"},
		//ProxyCountry:   "US",
		//ProfileId:      response.ProfileId, // Reuse the profileId data
		//ProfilePersist: true,               // Persist the browser session
	})

	log.Infof("%+v", browserInfo)
}
