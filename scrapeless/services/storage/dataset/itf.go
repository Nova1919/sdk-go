package dataset

import "context"

type Dataset interface {
	ListDatasets(ctx context.Context, page int64, pageSize int64, desc bool) (*ListDatasetsResponse, error)
	CreateDataset(ctx context.Context, name string) (id string, datasetName string, err error)
	UpdateDataset(ctx context.Context, datasetId string, name string) (ok bool, datasetName string, err error)
	DelDataset(ctx context.Context, datasetId string) (bool, error)
	AddItems(ctx context.Context, datasetId string, items []map[string]any) (bool, error)
	GetItems(ctx context.Context, datasetId string, page int, pageSize int, desc bool) (*ItemsResponse, error)

	Close() error
}
