package storage

import (
	"context"
	"github.com/smash-hq/sdk-go/env"
	"github.com/smash-hq/sdk-go/internal/code"
	"github.com/smash-hq/sdk-go/internal/remote/storage"
	"github.com/smash-hq/sdk-go/internal/remote/storage/models"
	"github.com/smash-hq/sdk-go/scrapeless/log"
)

type Dataset struct{}

// ListDatasets retrieves a list of dataset with pagination and sorting options.
// Parameters:
//
//	ctx: The request context.
//	page: Page number (starting from 1). Defaults to 1 if <=0.
//	pageSize:  Number of items per page. Minimum 10, defaults to 10 if smaller.
//	desc: Sort namespaces in descending order by creation time if true.
func (s *Dataset) ListDatasets(ctx context.Context, page int64, pageSize int64, desc bool) (*ListDatasetsResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 10 {
		pageSize = 10
	}
	datasets, err := storage.ClientInterface.ListDatasets(ctx, &models.ListDatasetsRequest{
		ActorId:  &env.GetActorEnv().ActorId,
		RunId:    &env.GetActorEnv().RunId,
		Page:     int(page),
		PageSize: int(pageSize),
		Desc:     desc,
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
		Total: int64(datasets.Total),
	}, nil
}

// CreateDataset Creates a new dataset storage.
// Parameters:
//
//	ctx:The request context.
//	name: The name of the dataset to create.
func (s *Dataset) CreateDataset(ctx context.Context, name string) (id string, datasetName string, err error) {
	name = name + "-" + env.GetActorEnv().RunId
	dataset, err := storage.ClientInterface.CreateDataset(ctx, &models.CreateDatasetRequest{
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
func (s *Dataset) UpdateDataset(ctx context.Context, datasetId string, name string) (ok bool, datasetName string, err error) {
	name = name + "-" + env.GetActorEnv().RunId
	ok, err = storage.ClientInterface.UpdateDataset(ctx, datasetId, name)
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
func (s *Dataset) DelDataset(ctx context.Context, datasetId string) (bool, error) {
	ok, err := storage.ClientInterface.DelDataset(ctx, datasetId)
	if err != nil {
		log.Errorf("failed to delete dataset: %v", code.Format(err))
		return false, code.Format(err)
	}
	return ok, nil
}

// AddItems adds a list of items to the dataset data store.
//
// Parameters:
//   - ctx: The context for the request.
//   - items: A slice of maps representing the items to add. Each map contains key-value pairs of any type.
func (s *Dataset) AddItems(ctx context.Context, datasetId string, items []map[string]any) (bool, error) {
	ok, err := storage.ClientInterface.AddDatasetItem(ctx, datasetId, items)
	if err != nil {
		log.Errorf("failed to add items: %v", err)
		return false, code.Format(err)
	}
	return ok, nil
}

// GetItems retrieves a list of items based on the provided pagination and sorting parameters.
//
// Parameters:
//
//	ctx: The context for the request.
//	page: The page number to retrieve (starting from 1).
//	pageSize: The number of items to return per page.
//	desc: Whether to sort items in descending order (true) or ascending (false).
func (s *Dataset) GetItems(ctx context.Context, datasetId string, page int, pageSize int, desc bool) (*ItemsResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 10 {
		pageSize = 10
	}
	items, err := storage.ClientInterface.GetDataset(ctx, &models.GetDataset{
		DatasetId: datasetId,
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

func (s *Dataset) Close() error {
	return nil
}
