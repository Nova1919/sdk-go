package kv

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/internal/code"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage/storage_http"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

type KvHttp struct {
	namespaceId string
}

func NewKVHttp(namespaceId ...string) KV {
	log.Infof("kv http init")
	if storage_http.Default() == nil {
		storage_http.Init(env.Env.ScrapelessStorageUrl)
	}
	kh := &KvHttp{namespaceId: env.GetActorEnv().KvNamespaceId}
	if len(namespaceId) > 0 {
		kh.namespaceId = namespaceId[0]
	}
	return kh
}

// ListNamespaces retrieves a list of KV namespaces with pagination and sorting options.
// Parameters:
//
//	ctx: The request context.
//	page: Page number (starting from 1). Defaults to 1 if <=0.
//	pageSize:  Number of items per page. Minimum 10, defaults to 10 if smaller.
//	desc: Sort namespaces in descending order by creation time if true.
func (kh *KvHttp) ListNamespaces(ctx context.Context, page int, pageSize int, desc bool) (*NamespacesResponse, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize < 10 {
		pageSize = 10
	}
	keyResp, err := storage_http.Default().ListNamespaces(ctx, page, pageSize, desc)
	if err != nil {
		log.Errorf("failed to list kv namespaces: %v", code.Format(err))
		return nil, code.Format(err)
	}
	var KvNamespaceItems []KvNamespaceItem
	for _, item := range keyResp.Items {
		namespaceItem := KvNamespaceItem{
			Id:         item.Id,
			Name:       item.Name,
			ActorId:    item.ActorId,
			RunId:      item.RunId,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
			AccessedAt: item.AccessedAt,
		}
		KvNamespaceItems = append(KvNamespaceItems, namespaceItem)
	}
	return &NamespacesResponse{
		Items: KvNamespaceItems,
		Total: keyResp.Total,
	}, nil
}

// CreateNamespace Creates a new key-value storage namespace.
// Parameters:
//
//	ctx:The request context.
//	name: The name of the namespace to create.
func (kh *KvHttp) CreateNamespace(ctx context.Context, name string) (namespaceId string, namespaceName string, err error) {
	name = name + "-" + env.GetActorEnv().RunId
	namespaceId, err = storage_http.Default().CreateNamespace(ctx, &storage_http.CreateKvNamespaceRequest{
		Name:    name,
		ActorId: env.GetActorEnv().ActorId,
		RunId:   env.GetActorEnv().RunId,
	})
	if err != nil {
		log.Errorf("failed to create kv namespace: %v", code.Format(err))
		return "", "", code.Format(err)
	}
	return namespaceId, name, nil
}

// GetNamespace retrieves namespace information by name
// Parameters:
//
//	ctx: The request context.
//	namespaceName: Name of the namespace to retrieve
func (kh *KvHttp) GetNamespace(ctx context.Context, namespaceName string) (*KvNamespaceItem, error) {
	namespace, err := storage_http.Default().GetNamespace(ctx, namespaceName)
	if err != nil {
		log.Errorf("failed to get kv namespace: %v", code.Format(err))
		return nil, code.Format(err)
	}
	resp := &KvNamespaceItem{
		Id:         namespace.Id,
		Name:       namespace.Name,
		ActorId:    namespace.ActorId,
		RunId:      namespace.RunId,
		CreatedAt:  namespace.CreatedAt,
		UpdatedAt:  namespace.UpdatedAt,
		AccessedAt: namespace.AccessedAt,
	}
	return resp, nil
}

// DelNamespace deletes the specified KV namespace.
// Parameters:
//
//	ctx:The request context.
func (kh *KvHttp) DelNamespace(ctx context.Context, namespaceId string) (bool, error) {
	ok, err := storage_http.Default().DelNamespace(ctx, namespaceId)
	if err != nil {
		log.Errorf("failed to delete kv namespace: %v", code.Format(err))
		return false, code.Format(err)
	}
	return ok, nil
}

// RenameNamespace renames an existing KV namespace
// Parameters:
//
//	ctx: The request context.
//	name: New namespace name
func (kh *KvHttp) RenameNamespace(ctx context.Context, namespaceId string, name string) (ok bool, namespaceName string, err error) {
	name = name + "-" + env.GetActorEnv().RunId
	ok, err = storage_http.Default().RenameNamespace(ctx, namespaceId, name)
	if err != nil {
		log.Errorf("failed to rename kv namespace: %v", code.Format(err))
		return false, "", code.Format(err)
	}
	return ok, name, nil
}

// ListKeys retrieves key list with pagination from the current namespace
// Parameters:
//
//	ctx: Request context
//	page: Page number (starting from 1). Defaults to 1 if <=0
//	pageSize: Number of items per page. Minimum 10, defaults to 10 if smaller
func (kh *KvHttp) ListKeys(ctx context.Context, namespaceId string, page int, pageSize int) (*KvKeys, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize < 10 {
		pageSize = 10
	}
	keys, err := storage_http.Default().ListKeys(ctx, &storage_http.ListKeyInfo{
		NamespaceId: namespaceId,
		Page:        page,
		Size:        pageSize,
	})
	if err != nil {
		log.Errorf("failed to list kv keys: %v", code.Format(err))
		return nil, code.Format(err)
	}
	if keys == nil {
		return nil, nil
	}
	kvKeys := &KvKeys{
		Items:     keys.Items,
		Total:     keys.Total,
		Page:      keys.Page,
		PageSize:  keys.PageSize,
		TotalPage: keys.TotalPage,
	}
	return kvKeys, nil
}
func (kh *KvHttp) DelValue(ctx context.Context, namespaceId string, key string) (bool, error) {
	ok, err := storage_http.Default().DelValue(ctx, namespaceId, key)
	if err != nil {
		log.Errorf("failed to delete kv value: %v", code.Format(err))
		return false, code.Format(err)
	}
	return ok, nil
}

func (kh *KvHttp) BulkSetValue(ctx context.Context, namespaceId string, data []BulkItem) (successCount int64, err error) {
	var items []storage_http.BulkItem
	for _, datum := range data {
		items = append(items, storage_http.BulkItem{
			Key:        datum.Key,
			Value:      datum.Value,
			Expiration: datum.Expiration,
		})
	}
	val, err := storage_http.Default().BulkSetValue(ctx, &storage_http.BulkSet{
		NamespaceId: namespaceId,
		Items:       items,
	})
	if err != nil {
		log.Errorf("failed to bulk set kv value: %v", code.Format(err))
		return 0, code.Format(err)
	}
	return val, nil
}
func (kh *KvHttp) BulkDelValue(ctx context.Context, namespaceId string, keys []string) (bool, error) {
	ok, err := storage_http.Default().BulkDelValue(ctx, namespaceId, keys)
	if err != nil {
		log.Errorf("failed to bulk delete kv value: %v", code.Format(err))
		return false, code.Format(err)
	}
	return ok, nil
}
func (kh *KvHttp) SetValue(ctx context.Context, namespaceId string, key string, value string, expiration uint) (bool, error) {
	ok, err := storage_http.Default().SetValue(ctx, &storage_http.SetValue{
		NamespaceId: namespaceId,
		Key:         key,
		Value:       value,
		Expiration:  expiration,
	})
	if err != nil {
		log.Errorf("failed to set kv value: %v", code.Format(err))
		return false, code.Format(err)
	}
	return ok, nil
}
func (kh *KvHttp) GetValue(ctx context.Context, namespaceId string, key string) (string, error) {
	val, err := storage_http.Default().GetValue(ctx, namespaceId, key)
	if err != nil {
		log.Errorf("failed to get kv value: %v", code.Format(err))
		return "", code.Format(err)
	}
	return val, nil
}
func (k *KvHttp) Close() error {
	return storage_http.Default().Close()
}
