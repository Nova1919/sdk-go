package main

import (
	"context"
	"github.com/smash-hq/sdk-go/scrapeless"
	"github.com/smash-hq/sdk-go/scrapeless/log"
)

func main() {
	client := scrapeless.New(scrapeless.WithExtension())

	// upload extension
	uploadExtension, err := client.Extension.Upload(context.Background(), "your-file-path.zip", "Scrapeless")
	if err != nil {
		panic(err)
	}
	log.Infof("Uploaded extension: %v", uploadExtension)

	// update extension
	updateExtensionSuccess, err := client.Extension.Update(context.Background(), uploadExtension.ExtensionID, "your-file-path.zip", "Scrapeless")
	if err != nil {
		panic(err)
	}
	log.Infof("Updated extension: %v", updateExtensionSuccess)

	// get extension
	extensionDetail, err := client.Extension.Get(context.Background(), "extensionID")
	if err != nil {
		panic(err)
	}
	log.Infof("Extension: %v", extensionDetail)

	// list extensions
	extensionList, err := client.Extension.List(context.Background())
	if err != nil {
		panic(err)
	}
	log.Infof("Extensions: %v", extensionList)

	// delete extension
	success, err := client.Extension.Delete(context.Background(), "extensionID")
	if err != nil {
		panic(err)
	}
	log.Infof("Deleted extension: %v", success)

}
