package main

import (
	"github.com/scrapeless-ai/sdk-go/scrapeless"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

func main() {
	client := scrapeless.New()
	defer client.Close()
	data, err := client.Router.Request("runnerId", "GET", "/v1/deepserp/search", nil, nil)
	if err != nil {
		panic(err)
	}
	log.Infof("%+v", data)
}
