package object

import (
	"context"
	"path/filepath"
	"strings"
)

type Object interface {
	ListBuckets(ctx context.Context, page int, pageSize int) (*ListBucketsResponse, error)
	CreateBucket(ctx context.Context, name string, description string) (bucketId string, bucketName string, err error)
	DeleteBucket(ctx context.Context) (bool, error)
	GetBucket(ctx context.Context) (*Bucket, error)
	List(ctx context.Context, fuzzyFileName string, page int64, pageSize int64) (*ListObjectsResponse, error)
	getWithId(ctx context.Context, objectId string) ([]byte, error)
	Get(ctx context.Context, objectId string) ([]byte, error)
	putWithId(ctx context.Context, filename string, data []byte) (string, error)
	Put(ctx context.Context, filename string, data []byte) (string, error)
	Delete(ctx context.Context, objectId string) (bool, error)
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
