package storage

import (
	"context"
	"errors"
	"github.com/smash-hq/sdk-go/env"
	"github.com/smash-hq/sdk-go/internal/code"
	"github.com/smash-hq/sdk-go/internal/remote/storage"
	"github.com/smash-hq/sdk-go/internal/remote/storage/models"
	"github.com/smash-hq/sdk-go/scrapeless/log"
	"path/filepath"
	"strings"
)

type Object struct{}

// ListBuckets retrieves the list of buckets with pagination support.
// Parameters:
//
//	ctx: The context for the request.
//	page: Current page number, minimum value is 1. Defaults to 1 if provided value is <1.
//	pageSize: Number of items per page, minimum value is 10. Defaults to 10 if provided value is <10.
func (s *Object) ListBuckets(ctx context.Context, page int, pageSize int) (*ListBucketsResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 10 {
		pageSize = 10
	}
	buckets, err := storage.ClientInterface.ListBuckets(ctx, page, pageSize)
	if err != nil {
		log.Errorf("failed to list buckets: %v", code.Format(err))
		return nil, code.Format(err)
	}
	var bucketsArray []Bucket
	for _, bucket := range buckets.Buckets {
		b := Bucket{
			Id:          bucket.Id,
			Name:        bucket.Name,
			Description: bucket.Description,
			CreatedAt:   bucket.CreatedAt,
			UpdatedAt:   bucket.UpdatedAt,
			ActorId:     bucket.ActorId,
			RunId:       bucket.RunId,
			Size:        bucket.Size,
		}
		bucketsArray = append(bucketsArray, b)
	}
	return &ListBucketsResponse{
		Buckets: bucketsArray,
		Total:   buckets.Total,
	}, nil
}

// CreateBucket creates a new storage bucket.
//
// Parameters:
//
//	ctx: The context for the request.
//	name: Bucket name, must comply with storage service naming rules.
//	description: Optional description for the bucket.
func (s *Object) CreateBucket(ctx context.Context, name string, description string) (bucketId string, bucketName string, err error) {
	name = name + "-" + env.GetActorEnv().RunId
	bucketId, err = storage.ClientInterface.CreateBucket(ctx, &models.CreateBucketRequest{
		Name:        name,
		Description: description,
		ActorId:     env.GetActorEnv().ActorId,
		RunId:       env.GetActorEnv().RunId,
	})
	if err != nil {
		log.Errorf("failed to create bucket: %v", code.Format(err))
		return "", "", code.Format(err)
	}
	return bucketId, bucketName, nil
}

// DeleteBucket delete bucket.
// Parameters:
//
//	ctx: The context for the request.
func (s *Object) DeleteBucket(ctx context.Context, bucketId string) (bool, error) {
	ok, err := storage.ClientInterface.DeleteBucket(ctx, bucketId)
	if err != nil {
		log.Errorf("failed to delete bucket: %v", code.Format(err))
		return false, code.Format(err)
	}
	return ok, nil
}

// GetBucket retrieves the bucket information associated with the ObjHttp instance.
// Parameters:
//
//	ctx: The context for the request.
func (s *Object) GetBucket(ctx context.Context, bucketId string) (*Bucket, error) {
	bucket, err := storage.ClientInterface.GetBucket(ctx, bucketId)
	if err != nil {
		log.Errorf("failed to get bucket: %v", code.Format(err))
		return nil, code.Format(err)
	}
	b := &Bucket{
		Id:          bucket.Id,
		Name:        bucket.Name,
		Description: bucket.Description,
		CreatedAt:   bucket.CreatedAt,
		UpdatedAt:   bucket.UpdatedAt,
		ActorId:     bucket.ActorId,
		RunId:       bucket.RunId,
		Size:        bucket.Size,
	}
	return b, nil
}

// List lists objects with fuzzy filename search and pagination support.
// Parameters:
//
//	ctx: The context for the request.
//	fuzzyFileName: Search pattern for matching object filenames.
//	page: Current page number, defaults to 1 if <1.
//	pageSize: Number of objects per page, defaults to 10 if <10.
func (s *Object) List(ctx context.Context, bucketId string, fuzzyFileName string, page int64, pageSize int64) (*ListObjectsResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 10 {
		pageSize = 10
	}
	objects, err := storage.ClientInterface.ListObjects(ctx, &models.ListObjectsRequest{
		BucketId: bucketId,
		Search:   fuzzyFileName,
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		log.Errorf("failed to list objects: %v", code.Format(err))
		return nil, code.Format(err)
	}
	var objectsArray []ObjectInfo
	for _, object := range objects.Objects {
		o := ObjectInfo{
			Id:        object.Id,
			Path:      object.Path,
			Size:      object.Size,
			Filename:  object.Filename,
			BucketId:  object.BucketId,
			ActorId:   object.ActorId,
			RunId:     object.RunId,
			FileType:  object.FileType,
			CreatedAt: object.CreatedAt,
			UpdatedAt: object.UpdatedAt,
		}
		objectsArray = append(objectsArray, o)
	}
	return &ListObjectsResponse{
		Objects: objectsArray,
		Total:   objects.Total,
	}, nil
}

// Get retrieves an object by its ID using HTTP.
//
// Parameters:
//
//	ctx: The context for the request.
//	objectId: The unique identifier of the object to retrieve.
func (s *Object) Get(ctx context.Context, bucketId string, objectId string) ([]byte, error) {
	object, err := storage.ClientInterface.GetObject(ctx, &models.ObjectRequest{
		BucketId: bucketId,
		ObjectId: objectId,
	})
	if err != nil {
		log.Errorf("failed to get object: %v", code.Format(err))
		return nil, code.Format(err)
	}
	return object, nil
}

// Put uploads the provided data to the object storage with the given filename.
//
// Parameters:
//
//	ctx: The context for the request.
//	filename: The name of the file to store.
//	data: The byte data to upload.
func (s *Object) Put(ctx context.Context, bucketId string, filename string, data []byte) (string, error) {
	_, ok := getObjectType(filename)
	if !ok {
		return "", errors.New("object type not supported")
	}
	object, err := storage.ClientInterface.PutObject(ctx, &models.PutObjectRequest{
		BucketId: bucketId,
		Filename: filename,
		Data:     data,
		ActorId:  env.GetActorEnv().ActorId,
		RunId:    env.GetActorEnv().RunId,
	})
	if err != nil {
		log.Errorf("failed to put object: %v", code.Format(err))
		return "", code.Format(err)
	}
	return object, nil
}

// Delete deletes an object from the specified bucket.
// Parameters:
//
//	ctx: The context used for the HTTP request.
//	objectId: The identifier of the object to delete.
func (s *Object) Delete(ctx context.Context, bucketId string, objectId string) (bool, error) {
	resp, err := storage.ClientInterface.DeleteObject(ctx, &models.ObjectRequest{
		BucketId: bucketId,
		ObjectId: objectId,
	})
	if err != nil {
		log.Errorf("failed to delete object: %v", code.Format(err))
		return false, code.Format(err)
	}
	return resp, nil
}

func (s *Object) Close() error {
	return nil
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
