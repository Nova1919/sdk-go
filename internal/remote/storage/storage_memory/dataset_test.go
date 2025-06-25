package storage_memory

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage/models"
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
	items, err := local.AddDatasetItem(ctx, "", maps)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(items)
}

func TestCreate(t *testing.T) {
	rep, err := local.CreateDataset(ctx, &models.CreateDatasetRequest{
		Name: "test-create",
	})
	if err != nil {
		t.Error(err)
	}
	marshal, _ := json.Marshal(rep)
	fmt.Println(string(marshal))
}

func TestDeleteDataset(t *testing.T) {
	dataset, err := local.DelDataset(ctx, datasetId)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(dataset)
}

func TestGetItems(t *testing.T) {
	items, err := local.GetDataset(ctx, &models.GetDataset{
		DatasetId: datasetId,
		Desc:      false,
		Page:      1,
		PageSize:  2,
	})
	if err != nil {
		t.Error(err)
	}
	marshal, _ := json.Marshal(items.Items)
	fmt.Println(string(marshal))
}

func TestListDataset(t *testing.T) {
	datasets, err := local.ListDatasets(ctx, &models.ListDatasetsRequest{
		ActorId:  nil,
		RunId:    nil,
		Page:     1,
		PageSize: 1,
		Desc:     false,
	})
	if err != nil {
		t.Error(err)
	}
	marshal, _ := json.Marshal(datasets)
	fmt.Println(string(marshal))
}

func TestUpdateDataset(t *testing.T) {
	ok, err := local.UpdateDataset(ctx, datasetId, "hq")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ok)
}
