package storage_memory

import (
	"context"
	"github.com/smash-hq/sdk-go/internal/remote/storage/models"
)

func (d *LocalClient) ListBuckets(ctx context.Context, page, size int) (*models.Object, error) {
	return nil, nil
}

func (d *LocalClient) CreateBucket(ctx context.Context, req *models.CreateBucketRequest) (string, error) {
	return "", nil
}

func (d *LocalClient) DeleteBucket(ctx context.Context, bucketId string) (bool, error) {
	return false, nil
}

func (d *LocalClient) GetBucket(ctx context.Context, bucketId string) (*models.Bucket, error) {
	return nil, nil
}

func (d *LocalClient) ListObjects(ctx context.Context, req *models.ListObjectsRequest) (*models.ObjectList, error) {
	return nil, nil
}

func (d *LocalClient) GetObject(ctx context.Context, req *models.ObjectRequest) ([]byte, error) {
	return nil, nil
}

func (d *LocalClient) DeleteObject(ctx context.Context, req *models.ObjectRequest) (bool, error) {
	return false, nil
}

func (d *LocalClient) PutObject(ctx context.Context, req *models.PutObjectRequest) (string, error) {
	return "", nil
}
