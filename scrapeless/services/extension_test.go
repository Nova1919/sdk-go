package services

import (
	"context"
	"github.com/smash-hq/sdk-go/scrapeless/services/extension"
	"testing"
)

func TestExtension(t *testing.T) {
	client := extension.NewExtension("http")

	list, err := client.List(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Log(list)

	//list, err := client.Upload(context.Background(), `D:\Backup\Downloads\大眼夹.zip`, "test")
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Log(list)

	//fmt.Println(client.Update(context.Background(), "79dfc808-df5b-44c2-aceb-94549913639e", `D:\Backup\Downloads\大眼夹.zip`, "test1"))
	//list, err := client.List(context.Background())
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Log(list)

	//detail, err := client.Get(context.Background(), "79dfc808-df5b-44c2-aceb-94549913639e")
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Log(detail)
	//
	//success, err := client.Delete(context.Background(), "79dfc808-df5b-44c2-aceb-94549913639e")
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Log(success)
}
