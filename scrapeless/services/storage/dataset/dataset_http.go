package dataset

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/internal/code"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage/storage_http"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

type DSHttp struct {
	datasetId string
}

func NewDSHttp(datasetId ...string) Dataset {
	log.Info("dataset http init")
	if storage_http.Default() == nil {
		storage_http.Init(env.Env.ScrapelessStorageUrl)
	}
	dh := &DSHttp{datasetId: env.GetActorEnv().DatasetId}
	if len(datasetId) > 0 {
		dh.datasetId = datasetId[0]
	}
	return dh
}

// ListDatasets retrieves a list of dataset with pagination and sorting options.
// Parameters:
//
//	ctx: The request context.
//	page: Page number (starting from 1). Defaults to 1 if <=0.
//	pageSize:  Number of items per page. Minimum 10, defaults to 10 if smaller.
//	desc: Sort namespaces in descending order by creation time if true.
func (ds *DSHttp) ListDatasets(ctx context.Context, page int64, pageSize int64, desc bool) (*ListDatasetsResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 10 {
		pageSize = 10
	}
	datasets, err := storage_http.Default().ListDatasets(ctx, &storage_http.ListDatasetsRequest{
		Page:     page,
		PageSize: pageSize,
		ActorId:  &env.GetActorEnv().ActorId,
		RunId:    &env.GetActorEnv().RunId,

		Desc: desc,
	})
	if err != nil {
		log.Errorf("failed to list datasets: %v", code.Format(err))
		return nil, code.Format(err)
	}
	var itemArray []DatasetInfo
	for _, item := range datasets.Items {
		itemArray = append(itemArray, DatasetInfo{
			Id:         item.Id,
			Name:       item.Name,
			ActorId:    item.ActorId,
			RunId:      item.RunId,
			Fields:     item.Fields,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
			AccessedAt: item.AccessedAt,
		})
	}
	return &ListDatasetsResponse{
		Items: itemArray,
		Total: datasets.Total,
	}, nil
}

// CreateDataset Creates a new dataset storage.
// Parameters:
//
//	ctx:The request context.
//	name: The name of the dataset to create.
func (ds *DSHttp) CreateDataset(ctx context.Context, name string) (id string, datasetName string, err error) {
	name = name + "-" + env.GetActorEnv().RunId
	dataset, err := storage_http.Default().CreateDataset(ctx, &storage_http.CreateDatasetRequest{
		Name:    name,
		ActorId: &env.GetActorEnv().ActorId,
		RunId:   &env.GetActorEnv().RunId,
	})
	if err != nil {
		log.Errorf("failed to create dataset: %v", code.Format(err))
		return "", "", code.Format(err)
	}
	return dataset.Id, name, nil
}

// UpdateDataset updates the dataset name by appending the current runtime ID to ensure uniqueness.
//
// Parameters:
//
//	ctx: The request context.
//	name: Original dataset name (will be combined with runtime ID internally)
func (ds *DSHttp) UpdateDataset(ctx context.Context, name string) (ok bool, datasetName string, err error) {
	name = name + "-" + env.GetActorEnv().RunId
	ok, err = storage_http.Default().UpdateDataset(ctx, ds.datasetId, name)
	if err != nil {
		log.Errorf("failed to update dataset: %v", code.Format(err))
		return false, "", code.Format(err)
	}
	return ok, name, nil
}

// DelDataset deletes a dataset asynchronously.
//
// Parameters:
//
//	ctx: The context for the request, used for cancellation and timeouts.
func (ds *DSHttp) DelDataset(ctx context.Context) (bool, error) {
	ok, err := storage_http.Default().DelDataset(ctx, ds.datasetId)
	if err != nil {
		log.Errorf("failed to delete dataset: %v", code.Format(err))
		return false, code.Format(err)
	}
	return ok, nil
}

func (ds *DSHttp) addItemsWithId(ctx context.Context, items []map[string]any) (bool, error) {
	ok, err := storage_http.Default().AddDatasetItem(ctx, ds.datasetId, items)
	if err != nil {
		log.Errorf("failed to add items: %v", err)
		return false, code.Format(err)
	}
	return ok, nil
}

func (ds *DSHttp) getItemsWithId(ctx context.Context, page int, pageSize int, desc bool) (*ItemsResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 10 {
		pageSize = 10
	}
	items, err := storage_http.Default().GetDataset(ctx, &storage_http.GetDataset{
		DatasetId: ds.datasetId,
		Desc:      desc,
		Page:      page,
		PageSize:  pageSize,
	})
	if err != nil {
		log.Errorf("failed to get items: %v", code.Format(err))
		return nil, code.Format(err)
	}
	var itemArray []map[string]any
	for _, item := range items.Items {
		itemArray = append(itemArray, item)
	}
	return &ItemsResponse{
		Items: itemArray,
		Total: items.Total,
	}, nil
}

// AddItems adds a list of items to the dataset data store.
//
// Parameters:
//   - ctx: The context for the request.
//   - items: A slice of maps representing the items to add. Each map contains key-value pairs of any type.
func (ds *DSHttp) AddItems(ctx context.Context, items []map[string]any) (bool, error) {
	return ds.addItemsWithId(ctx, items)
}

// GetItems retrieves a list of items based on the provided pagination and sorting parameters.
//
// Parameters:
//
//	ctx: The context for the request.
//	page: The page number to retrieve (starting from 1).
//	pageSize: The number of items to return per page.
//	desc: Whether to sort items in descending order (true) or ascending (false).
func (ds *DSHttp) GetItems(ctx context.Context, page int, pageSize int, desc bool) (*ItemsResponse, error) {
	return ds.getItemsWithId(ctx, page, pageSize, desc)
}

func (ds *DSHttp) Close() error {
	return storage_http.Default().Close()
}
