package storage_memory

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func init() {
	Init()
}

var (
	datasetId = "123456"
	local     = DatasetLocal{datasetId: datasetId}
	ctx       = context.Background()
)

func TestAddDataset(t *testing.T) {
	maps := []map[string]interface{}{
		{"name": "hq", "sex": "man", "age": "18"},
		{"name": "wu", "sex": "woman", "age": "28"},
		{"name": "op", "sex": "man", "age": "38"},
	}
	items, err := local.AddItems(ctx, "", maps)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(items)
}

func TestDeleteDataset(t *testing.T) {
	dataset, err := local.DelDataset(ctx, datasetId)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(dataset)
}

func TestGetItems(t *testing.T) {
	items, err := local.GetItems(ctx, datasetId, 1, 10, true)
	if err != nil {
		t.Error(err)
	}
	marshal, _ := json.Marshal(items.Items)
	fmt.Println(string(marshal))
}

func TestListDataset(t *testing.T) {
	datasets, err := local.ListDatasets(ctx, 1, 10, true)
	if err != nil {
		t.Error(err)
	}
	marshal, _ := json.Marshal(datasets)
	fmt.Println(string(marshal))
}

func TestUpdateDataset(t *testing.T) {
	ok, name, err := local.UpdateDataset(ctx, datasetId, "hq")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ok, name)
}
