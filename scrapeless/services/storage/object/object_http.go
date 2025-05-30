package object

import (
	"context"
	"errors"
	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/internal/code"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage/storage_http"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

type ObjHttp struct {
	bucketId string
}

func NewObjHttp(bucketId ...string) Object {
	log.Info("object http init")
	if storage_http.Default() == nil {
		storage_http.Init(env.Env.ScrapelessStorageUrl)
	}
	oh := &ObjHttp{bucketId: env.GetActorEnv().BucketId}
	if len(bucketId) > 0 {
		oh.bucketId = bucketId[0]
	}
	return oh
}

// ListBuckets retrieves the list of buckets with pagination support.
// Parameters:
//
//	ctx: The context for the request.
//	page: Current page number, minimum value is 1. Defaults to 1 if provided value is <1.
//	pageSize: Number of items per page, minimum value is 10. Defaults to 10 if provided value is <10.
func (oh *ObjHttp) ListBuckets(ctx context.Context, page int, pageSize int) (*ListBucketsResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 10 {
		pageSize = 10
	}
	buckets, err := storage_http.Default().ListBuckets(ctx, page, pageSize)
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
func (oh *ObjHttp) CreateBucket(ctx context.Context, name string, description string) (bucketId string, bucketName string, err error) {
	name = name + "-" + env.GetActorEnv().RunId
	bucketId, err = storage_http.Default().CreateBucket(ctx, &storage_http.CreateBucketRequest{
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
func (oh *ObjHttp) DeleteBucket(ctx context.Context) (bool, error) {
	ok, err := storage_http.Default().DeleteBucket(ctx, oh.bucketId)
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
func (oh *ObjHttp) GetBucket(ctx context.Context) (*Bucket, error) {
	bucket, err := storage_http.Default().GetBucket(ctx, oh.bucketId)
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
func (oh *ObjHttp) List(ctx context.Context, fuzzyFileName string, page int64, pageSize int64) (*ListObjectsResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 10 {
		pageSize = 10
	}
	objects, err := storage_http.Default().ListObjects(ctx, &storage_http.ListObjectsRequest{
		BucketId: oh.bucketId,
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

func (oh *ObjHttp) getWithId(ctx context.Context, objectId string) ([]byte, error) {
	object, err := storage_http.Default().GetObject(ctx, &storage_http.ObjectRequest{
		BucketId: oh.bucketId,
		ObjectId: objectId,
	})
	if err != nil {
		log.Errorf("failed to get object: %v", code.Format(err))
		return nil, code.Format(err)
	}
	return object, nil
}

// Get retrieves an object by its ID using HTTP.
//
// Parameters:
//
//	ctx: The context for the request.
//	objectId: The unique identifier of the object to retrieve.
func (oh *ObjHttp) Get(ctx context.Context, objectId string) ([]byte, error) {
	return oh.getWithId(ctx, objectId)
}

func (oh *ObjHttp) putWithId(ctx context.Context, filename string, data []byte) (string, error) {
	_, ok := getObjectType(filename)
	if !ok {
		return "", errors.New("object type not supported")
	}
	object, err := storage_http.Default().PutObject(ctx, &storage_http.PutObjectRequest{
		BucketId: oh.bucketId,
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

// Put uploads the provided data to the object storage with the given filename.
//
// Parameters:
//
//	ctx: The context for the request.
//	filename: The name of the file to store.
//	data: The byte data to upload.
func (oh *ObjHttp) Put(ctx context.Context, filename string, data []byte) (string, error) {
	return oh.putWithId(ctx, filename, data)
}

// Delete deletes an object from the specified bucket.
// Parameters:
//
//	ctx: The context used for the HTTP request.
//	objectId: The identifier of the object to delete.
func (oh *ObjHttp) Delete(ctx context.Context, objectId string) (bool, error) {
	resp, err := storage_http.Default().DeleteObject(ctx, &storage_http.ObjectRequest{
		BucketId: oh.bucketId,
		ObjectId: objectId,
	})
	if err != nil {
		log.Errorf("failed to delete object: %v", code.Format(err))
		return false, code.Format(err)
	}
	return resp, nil
}

func (oh *ObjHttp) Close() error {
	return storage_http.Default().Close()
}
