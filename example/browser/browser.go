package main

import (
	"context"
	"github.com/smash-hq/sdk-go/scrapeless"
	"github.com/smash-hq/sdk-go/scrapeless/log"
	"github.com/smash-hq/sdk-go/scrapeless/services/browser"
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

	// upload extension
	uploadExtension, err := client.Browser.Upload(context.Background(), "your-file-path.zip", "Scrapeless")
	if err != nil {
		panic(err)
	}
	log.Infof("Uploaded extension: %v", uploadExtension)

	// update extension
	updateExtensionSuccess, err := client.Browser.Update(context.Background(), uploadExtension.ExtensionID, "your-file-path.zip", "Scrapeless")
	if err != nil {
		panic(err)
	}
	log.Infof("Updated extension: %v", updateExtensionSuccess)

	// get extension
	extensionDetail, err := client.Browser.Get(context.Background(), "extensionID")
	if err != nil {
		panic(err)
	}
	log.Infof("Extension: %v", extensionDetail)

	// list extensions
	extensionList, err := client.Browser.List(context.Background())
	if err != nil {
		panic(err)
	}
	log.Infof("Extensions: %v", extensionList)

	// delete extension
	success, err := client.Browser.Delete(context.Background(), "extensionID")
	if err != nil {
		panic(err)
	}
	log.Infof("Deleted extension: %v", success)
}
