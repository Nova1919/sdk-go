package object

import (
	"context"
	"path/filepath"
	"strings"
)

type Object interface {
	ListBuckets(ctx context.Context, page int, pageSize int) (*ListBucketsResponse, error)
	CreateBucket(ctx context.Context, name string, description string) (bucketId string, bucketName string, err error)
	DeleteBucket(ctx context.Context, bucketId string) (bool, error)
	GetBucket(ctx context.Context, bucketId string) (*Bucket, error)
	List(ctx context.Context, bucketId string, fuzzyFileName string, page int64, pageSize int64) (*ListObjectsResponse, error)
	Get(ctx context.Context, bucketId string, objectId string) ([]byte, error)
	Put(ctx context.Context, bucketId string, filename string, data []byte) (string, error)
	Delete(ctx context.Context, bucketId string, objectId string) (bool, error)
	Close() error
}

var (
	ObjectTypeMapping = map[string]struct{}{
		"json": {},
		"html": {},
		"png":  {},
	}
)

func getObjectType(filename string) (string, bool) {
	ext := filepath.Ext(filename)
	ext = strings.Replace(ext, ".", "", -1)
	_, ok := ObjectTypeMapping[ext]
	return ext, ok
}
