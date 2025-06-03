package kv

import (
	"context"
)

type KV interface {
	ListNamespaces(ctx context.Context, page int, pageSize int, desc bool) (*NamespacesResponse, error)
	CreateNamespace(ctx context.Context, name string) (namespaceId string, namespaceName string, err error)
	GetNamespace(ctx context.Context, namespaceName string) (*KvNamespaceItem, error)
	DelNamespace(ctx context.Context, namespaceId string) (bool, error)
	RenameNamespace(ctx context.Context, namespaceId string, name string) (ok bool, namespaceName string, err error)
	ListKeys(ctx context.Context, namespaceId string, page int, pageSize int) (*KvKeys, error)
	SetValue(ctx context.Context, namespaceId string, key string, value string, expiration uint) (bool, error)
	DelValue(ctx context.Context, namespaceId string, key string) (bool, error)
	BulkSetValue(ctx context.Context, namespaceId string, data []BulkItem) (successCount int64, err error)
	BulkDelValue(ctx context.Context, namespaceId string, keys []string) (bool, error)
	GetValue(ctx context.Context, namespaceId string, key string) (string, error)

	Close() error
}
