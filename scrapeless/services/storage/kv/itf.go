package kv

import (
	"context"
)

type KV interface {
	ListNamespaces(ctx context.Context, page int, pageSize int, desc bool) (*NamespacesResponse, error)
	CreateNamespace(ctx context.Context, name string) (namespaceId string, namespaceName string, err error)
	GetNamespace(ctx context.Context, namespaceName string) (*KvNamespaceItem, error)
	DelNamespace(ctx context.Context) (bool, error)
	RenameNamespace(ctx context.Context, name string) (ok bool, namespaceName string, err error)
	ListKeys(ctx context.Context, page int, pageSize int) (*KvKeys, error)
	setValueWithId(ctx context.Context, key string, value string, expiration uint) (bool, error)
	DelValue(ctx context.Context, key string) (bool, error)
	getValueWithId(ctx context.Context, key string) (string, error)
	BulkSetValue(ctx context.Context, data []BulkItem) (successCount int64, err error)
	BulkDelValue(ctx context.Context, keys []string) (bool, error)
	SetValue(ctx context.Context, key string, value string, expiration uint) (bool, error)
	GetValue(ctx context.Context, key string) (string, error)

	Close() error
}
