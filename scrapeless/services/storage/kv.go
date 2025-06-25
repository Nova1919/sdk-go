package storage

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/internal/code"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage/models"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

type KV struct{}

// ListNamespaces retrieves a list of KV namespaces with pagination and sorting options.
// Parameters:
//
//	ctx: The request context.
//	page: Page number (starting from 1). Defaults to 1 if <=0.
//	pageSize:  Number of items per page. Minimum 10, defaults to 10 if smaller.
//	desc: Sort namespaces in descending order by creation time if true.
func (s *KV) ListNamespaces(ctx context.Context, page int64, pageSize int64, desc bool) (*NamespacesResponse, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize < 10 {
		pageSize = 10
	}
	keyResp, err := storage.ClientInterface.ListNamespaces(ctx, page, pageSize, desc)
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
func (s *KV) CreateNamespace(ctx context.Context, name string) (namespaceId string, namespaceName string, err error) {
	name = name + "-" + env.GetActorEnv().RunId
	namespaceId, err = storage.ClientInterface.CreateNamespace(ctx, &models.CreateKvNamespaceRequest{
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
func (s *KV) GetNamespace(ctx context.Context, namespaceName string) (*KvNamespaceItem, error) {
	namespace, err := storage.ClientInterface.GetNamespace(ctx, namespaceName)
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
func (s *KV) DelNamespace(ctx context.Context, namespaceId string) (bool, error) {
	ok, err := storage.ClientInterface.DelNamespace(ctx, namespaceId)
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
func (s *KV) RenameNamespace(ctx context.Context, namespaceId string, name string) (ok bool, namespaceName string, err error) {
	name = name + "-" + env.GetActorEnv().RunId
	ok, err = storage.ClientInterface.RenameNamespace(ctx, namespaceId, name)
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
func (s *KV) ListKeys(ctx context.Context, namespaceId string, page int, pageSize int) (*KvKeys, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize < 10 {
		pageSize = 10
	}
	keys, err := storage.ClientInterface.ListKeys(ctx, &models.ListKeyInfo{
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

// DelValue deletes the value associated with the specified key in the given namespace.
// Parameters:
//
//	ctx: Request context
//	namespaceId: Identifier of the namespace
//	key: The key to delete
func (s *KV) DelValue(ctx context.Context, namespaceId string, key string) (bool, error) {
	ok, err := storage.ClientInterface.DelValue(ctx, namespaceId, key)
	if err != nil {
		log.Errorf("failed to delete kv value: %v", code.Format(err))
		return false, code.Format(err)
	}
	return ok, nil
}

// BulkSetValue sets multiple key-value pairs in the specified namespace.
// Parameters:
//
//	ctx: Request context
//	namespaceId: Identifier of the namespace
//	data: A slice of BulkItem containing key, value, and optional expiration
func (s *KV) BulkSetValue(ctx context.Context, namespaceId string, data []BulkItem) (successCount int64, err error) {
	var items []models.BulkItem
	for _, datum := range data {
		items = append(items, models.BulkItem{
			Key:        datum.Key,
			Value:      datum.Value,
			Expiration: datum.Expiration,
		})
	}

	val, err := storage.ClientInterface.BulkSetValue(ctx, &models.BulkSet{
		NamespaceId: namespaceId,
		Items:       items,
	})
	if err != nil {
		log.Errorf("failed to bulk set kv value: %v", code.Format(err))
		return 0, code.Format(err)
	}
	return val, nil
}

// BulkDelValue deletes multiple keys from the specified namespace.
// Parameters:
//
//	ctx: Request context
//	namespaceId: Identifier of the namespace
//	keys: A slice of keys to delete
func (s *KV) BulkDelValue(ctx context.Context, namespaceId string, keys []string) (bool, error) {
	ok, err := storage.ClientInterface.BulkDelValue(ctx, namespaceId, keys)
	if err != nil {
		log.Errorf("failed to bulk delete kv value: %v", code.Format(err))
		return false, code.Format(err)
	}
	return ok, nil
}

// SetValue sets a key-value pair in the specified namespace.
// Parameters:
//
//	ctx: Request context
//	namespaceId:
//	key: kv key
//	value: kv value
//	expiration: kv expiration  Time-to-live in seconds (s)
func (s *KV) SetValue(ctx context.Context, namespaceId string, key string, value string, expiration uint) (bool, error) {
	ok, err := storage.ClientInterface.SetValue(ctx, &models.SetValue{
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

// GetValue retrieves the value for the specified key in the given namespace.
// Parameters:
//
//	ctx: Request context
//	namespaceId: Identifier of the namespace
//	key: The key whose value is to be retrieved
func (s *KV) GetValue(ctx context.Context, namespaceId string, key string) (string, error) {
	val, err := storage.ClientInterface.GetValue(ctx, namespaceId, key)
	if err != nil {
		log.Errorf("failed to get kv value: %v", code.Format(err))
		return "", code.Format(err)
	}
	return val, nil
}

func (s *KV) Close() error {
	return nil
}
