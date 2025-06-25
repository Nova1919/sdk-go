package main

import (
	"github.com/smash-hq/sdk-go/scrapeless"
	sh "github.com/smash-hq/sdk-go/scrapeless/services/httpserver"
)

func main() {
	client := scrapeless.New(scrapeless.WithServer())
	defer client.Close()

	// input is map[string]string
	client.Server.AddHandleGet("/", func(input []byte) (sh.Response, error) {
		return sh.Response{
			Code: 200,
			Data: "Hello World",
		}, nil
	})

	// input is json
	client.Server.AddHandlePost("/", func(input []byte) (sh.Response, error) {
		return sh.Response{
			Code: 200,
			Data: "Hello World",
		}, nil
	})

	_ = client.Server.Start()
}
