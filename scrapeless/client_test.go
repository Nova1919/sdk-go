package scrapeless

import (
	"context"
	"testing"
)

func TestNew(t *testing.T) {
	client := New(WithStorage())
	namespaces, err := client.Storage.Kv.ListNamespaces(context.Background(), 1, 10, true)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("%v", namespaces)
}
