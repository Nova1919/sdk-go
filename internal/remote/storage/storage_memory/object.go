package storage_memory

import (
	"context"
	"github.com/smash-hq/sdk-go/internal/remote/storage/models"
)

func (c *LocalClient) ListBuckets(ctx context.Context, page, size int) (*models.Object, error) {
	return nil, nil
}

func (c *LocalClient) CreateBucket(ctx context.Context, req *models.CreateBucketRequest) (string, error) {
	return "", nil
}

func (c *LocalClient) DeleteBucket(ctx context.Context, bucketId string) (bool, error) {
	return false, nil
}

func (c *LocalClient) GetBucket(ctx context.Context, bucketId string) (*models.Bucket, error) {
	return nil, nil
}

func (c *LocalClient) ListObjects(ctx context.Context, req *models.ListObjectsRequest) (*models.ObjectList, error) {
	return nil, nil
}

func (c *LocalClient) GetObject(ctx context.Context, req *models.ObjectRequest) ([]byte, error) {
	return nil, nil
}

func (c *LocalClient) DeleteObject(ctx context.Context, req *models.ObjectRequest) (bool, error) {
	return false, nil
}

func (c *LocalClient) PutObject(ctx context.Context, req *models.PutObjectRequest) (string, error) {
	return "", nil
}
