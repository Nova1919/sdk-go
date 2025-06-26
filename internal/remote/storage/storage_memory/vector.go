package storage_memory

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage/models"
)

func (c *LocalClient) ListCollections(ctx context.Context, req *models.ListCollectionsRequest) (*models.ListCollectionsResponse, error) {
	return nil, nil
}

func (c *LocalClient) CreateCollections(ctx context.Context, req *models.CreateCollectionRequest) (*models.CreateCollectionResponse, error) {
	return nil, nil
}

func (c *LocalClient) UpdateCollection(ctx context.Context, req *models.UpdateCollectionRequest) error {
	return nil
}

func (c *LocalClient) DelCollection(ctx context.Context, collId string) error {
	return nil
}

func (c *LocalClient) GetCollection(ctx context.Context, collId string) (*models.Collection, error) {
	return nil, nil
}

func (c *LocalClient) CreateDocs(ctx context.Context, req *models.CreateDocsRequest) (*models.DocOpResponse, error) {
	return nil, nil
}

func (c *LocalClient) UpdateDocs(ctx context.Context, req *models.UpdateDocsRequest) (*models.DocOpResponse, error) {
	return nil, nil
}

func (c *LocalClient) UpsertDocs(ctx context.Context, req *models.UpsertVectorDocsParam) (*models.DocOpResponse, error) {
	return nil, nil
}

func (c *LocalClient) DelDocs(ctx context.Context, req *models.DeleteDocsRequest) (*models.DocOpResponse, error) {
	return nil, nil
}

func (c *LocalClient) QueryDocs(ctx context.Context, req *models.QueryVectorRequest) ([]*models.Doc, error) {
	return nil, nil
}

func (c *LocalClient) QueryDocsByIds(ctx context.Context, req *models.QueryDocsByIdsRequest) (map[string]*models.Doc, error) {
	return nil, nil
}
