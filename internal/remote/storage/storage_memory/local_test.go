package storage_memory

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/smash-hq/sdk-go/internal/remote/storage/models"
	"testing"
)

func init() {
	Init()
}

var (
	datasetId   = "123456"
	NamespaceId = "1245434234"
	local       = DatasetLocal{datasetId: datasetId}
	kvLocal     = KVLocal{NamespaceId: NamespaceId}
	queueLocal  = QueueLocal{}
	ctx         = context.Background()
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

func TestCreateNamespace(t *testing.T) {

	Id, err := kvLocal.CreateNamespace(ctx, &models.CreateKvNamespaceRequest{
		Name:    "hq11343342",
		RunId:   "hq11",
		ActorId: "hq11",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(Id, err)
}

func TestListNamespace(t *testing.T) {

	ns, err := kvLocal.ListNamespaces(ctx, 2, 2, true)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ns, err)
}

func TestDelNamespace(t *testing.T) {

	ns, err := kvLocal.DelNamespace(ctx, "6644afa6-4c9e-4904-97ce-4cc30f2e75ff")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ns, err)
}

func TestGetNamespace(t *testing.T) {
	ns, err := kvLocal.GetNamespace(ctx, "56af1a69-8a9f-44eb-8b05-b251a5122619")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ns, err)
}

func TestRenameNamespace(t *testing.T) {
	ns, err := kvLocal.RenameNamespace(ctx, "56af1a69-8a9f-44eb-8b05-b251a5122619", "bbbbsd")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ns, err)
}

func TestSetValue(t *testing.T) {
	ok, err := kvLocal.SetValue(ctx, &models.SetValue{
		NamespaceId: "56af1a69-8a9f-44eb-8b05-b251a5122619",
		Key:         "Mykey5",
		Value:       "myValue",
		Expiration:  10,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ok, err)
}

func TestListKeys(t *testing.T) {
	ok, err := kvLocal.ListKeys(ctx, &models.ListKeyInfo{
		NamespaceId: "56af1a69-8a9f-44eb-8b05-b251a5122619",
		Page:        1,
		Size:        9,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ok, err)
}

func TestBulkSetValue(t *testing.T) {
	ok, err := kvLocal.BulkSetValue(ctx, &models.BulkSet{
		NamespaceId: "56af1a69-8a9f-44eb-8b05-b251a5122619",
		Items: []models.BulkItem{
			{Key: "Mykey5", Value: "myValue"},
			{Key: "Mykey6", Value: "myValue"},
			{Key: "Mykey7", Value: "myValue"},
			{Key: "Mykey8", Value: "myValue"},
		},
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ok, err)
}

func TestBulkDelValue(t *testing.T) {
	ok, err := kvLocal.BulkDelValue(ctx, "56af1a69-8a9f-44eb-8b05-b251a5122619", []string{"Mykey7", "Mykey8"})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ok, err)
}

func TestDelValue(t *testing.T) {
	ok, err := kvLocal.DelValue(ctx, "56af1a69-8a9f-44eb-8b05-b251a5122619", "metadata")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ok, err)
}

func TestGetValue(t *testing.T) {
	ok, err := kvLocal.GetValue(ctx, "56af1a69-8a9f-44eb-8b05-b251a5122619", "Mykey2")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ok, err)
}

func TestCreateQueue(t *testing.T) {
	ok, err := queueLocal.CreateQueue(ctx, &models.CreateQueueRequest{
		Name:        "myQuefwudwae23wrq2",
		Description: "dwafmaio",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ok, err)
}

func TestGetQueue(t *testing.T) {
	ok, err := queueLocal.GetQueue(ctx, &models.GetQueueRequest{
		Id: "2050499a-f190-4d03-a1cc-335a438c5c6d",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ok, err)
}

func TestListQueue(t *testing.T) {
	ok, err := queueLocal.GetQueues(ctx, &models.GetQueuesRequest{
		Page:     1,
		PageSize: 3,
		Desc:     false,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ok, err)
}

func TestUpdateQueue(t *testing.T) {
	err := queueLocal.UpdateQueue(ctx, &models.UpdateQueueRequest{
		QueueId:     "2050499a-f190-4d03-a1cc-335a438c5c6d",
		Name:        "6666",
		Description: "myQueue",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestCreateMsg(t *testing.T) {
	ok, err := queueLocal.CreateMsg(ctx, &models.CreateMsgRequest{
		QueueId:  "2050499a-f190-4d03-a1cc-335a438c5c6d",
		Name:     "6666",
		PayLoad:  "myQueue",
		Retry:    3,
		Timeout:  10,
		Deadline: 1750836277,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ok, err)
}
