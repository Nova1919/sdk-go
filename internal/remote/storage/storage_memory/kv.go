package storage_memory

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage/models"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"time"
)

var (
	ErrResourceNotFound = errors.New("resource not found")
	ErrResourceExists   = errors.New("resource exists")
)

const MaxExpireTime = 24 * 60 * 60 * 7

func (d *LocalClient) GetNamespace(ctx context.Context, namespaceId string) (*models.KvNamespaceItem, error) {
	nsPath := filepath.Join(storageDir, keyValueDir, namespaceId)
	ok := isDirExists(nsPath)
	if !ok {
		return nil, ErrResourceNotFound
	}
	metaDataPath := filepath.Join(nsPath, metadataFile)
	file, err := os.ReadFile(metaDataPath)
	if err != nil {
		return nil, fmt.Errorf("read file %s failed: %v", metaDataPath, err)
	}

	var namespace models.KvNamespaceItem
	if err = json.Unmarshal(file, &namespace); err != nil {
		return nil, fmt.Errorf("json unmarshal failed: %s", err)
	}

	return &namespace, nil
}

func (d *LocalClient) ListNamespaces(ctx context.Context, page int64, pageSize int64, desc bool) (*models.KvNamespace, error) {
	dirPath := filepath.Join(storageDir, keyValueDir)

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir: %v", err)
	}

	var allNamespaces []models.KvNamespaceItem

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		name := entry.Name()
		metaPath := filepath.Join(dirPath, name, metadataFile)

		file, err := os.ReadFile(metaPath)
		if err != nil {
			continue
		}

		var meta models.KvNamespaceItem
		if err = json.Unmarshal(file, &meta); err != nil {
			continue
		}

		allNamespaces = append(allNamespaces, meta)
	}

	// sort
	sort.Slice(allNamespaces, func(i, j int) bool {
		if desc {
			return allNamespaces[i].CreatedAt > allNamespaces[j].CreatedAt
		}
		return allNamespaces[i].CreatedAt < allNamespaces[j].CreatedAt
	})

	total := int64(len(allNamespaces))

	// page
	start := (page - 1) * pageSize
	if start > total {
		start = total
	}
	end := start + pageSize
	if end > total {
		end = total
	}

	pagedItems := allNamespaces[start:end]

	return &models.KvNamespace{
		Items:     pagedItems,
		Total:     total,
		Page:      page,
		PageSize:  pageSize,
		TotalPage: totalPage(total, pageSize),
	}, nil
}

func (d *LocalClient) CreateNamespace(ctx context.Context, req *models.CreateKvNamespaceRequest) (namespaceId string, err error) {
	id := uuid.NewString()
	path := filepath.Join(storageDir, keyValueDir, id)

	exists, err := isNameExists(filepath.Join(storageDir, keyValueDir), req.Name)
	if err != nil {
		return "", err
	}
	if exists {
		return "", fmt.Errorf("namespace %s already exists", req.Name)
	}

	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("create dataset failed, cause: %v", err)
	}

	namespace := models.KvNamespaceItem{
		Id:        id,
		Name:      req.Name,
		RunId:     req.RunId,
		ActorId:   req.ActorId,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
	marshal, err := json.Marshal(&namespace)
	if err != nil {
		return "", fmt.Errorf("marshal namespace failed, cause: %v", err)
	}
	metaFile := filepath.Join(path, metadataFile)

	err = os.WriteFile(metaFile, marshal, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("write file %s failed, cause: %v", metaFile, err)
	}
	return id, nil
}

func (d *LocalClient) DelNamespace(ctx context.Context, namespaceId string) (bool, error) {
	absPath := filepath.Join(storageDir, keyValueDir, namespaceId)
	err := os.RemoveAll(absPath)
	if err != nil {
		return false, fmt.Errorf("delete namespace failed, cause: %v", err)
	}
	return true, nil
}

func (d *LocalClient) RenameNamespace(ctx context.Context, namespaceId string, name string) (ok bool, err error) {
	nsPath := filepath.Join(storageDir, keyValueDir, namespaceId)
	exists := isDirExists(nsPath)
	if !exists {
		return false, ErrResourceNotFound
	}
	filePath := filepath.Join(nsPath, metadataFile)
	file, err := os.ReadFile(filePath)
	if err != nil {
		return false, fmt.Errorf("read file %s failed: %v", filePath, err)
	}

	var old models.KvNamespaceItem
	if err = json.Unmarshal(file, &old); err != nil {
		return false, fmt.Errorf("json unmarshal failed: %s", err)
	}
	old.Name = name

	marshal, err := json.Marshal(&old)
	if err != nil {
		return false, fmt.Errorf("json marshal failed: %s", err)
	}

	err = os.WriteFile(filePath, marshal, 0644)
	if err != nil {
		return false, fmt.Errorf("write file %s failed: %v", filePath, err)
	}
	return true, nil
}

func (d *LocalClient) SetValue(ctx context.Context, req *models.SetValue) (bool, error) {
	path := filepath.Join(storageDir, keyValueDir, req.NamespaceId)
	file := filepath.Join(path, fmt.Sprintf("%s.json", req.Key))
	if req.Expiration == 0 {
		req.Expiration = MaxExpireTime
	}
	local := models.SetValueLocal{
		SetValue: models.SetValue{
			Expiration:  req.Expiration,
			Key:         req.Key,
			Value:       req.Value,
			NamespaceId: req.NamespaceId,
		},
		ExpireAt: time.Now().Add(time.Duration(req.Expiration) * time.Second),
		Size:     len([]byte(req.Value)),
	}

	marshal, err := json.Marshal(local)
	if err != nil {
		return false, fmt.Errorf("json marshal failed: %s", err)
	}
	err = os.WriteFile(file, marshal, os.ModePerm)

	return true, nil
}

func (d *LocalClient) ListKeys(ctx context.Context, req *models.ListKeyInfo) (*models.KvKeys, error) {
	dirPath := filepath.Join(storageDir, keyValueDir, req.NamespaceId)
	var keys []map[string]any

	now := time.Now()
	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || d.Name() == metadataFile {
			return nil
		}

		kvFile, err := os.ReadFile(filepath.Join(dirPath, d.Name()))
		if err != nil {
			return fmt.Errorf("read file %s failed: %v", path, err)
		}
		var kv models.SetValueLocal
		err = json.Unmarshal(kvFile, &kv)
		if err != nil {
			return fmt.Errorf("json unmarshal failed: %s", err)
		}

		if kv.ExpireAt.Before(now) {
			return nil
		}

		keys = append(keys, map[string]any{
			"key":  kv.Key,
			"size": kv.Size,
		})

		return nil
	})
	if err != nil {
		return nil, err
	}
	total := int64(len(keys))
	kvKeys := &models.KvKeys{
		Total:     total,
		Page:      req.Page,
		PageSize:  req.Size,
		TotalPage: totalPage(total, req.Size),
	}

	start := (req.Page - 1) * req.Size
	if start >= total {
		return kvKeys, nil
	}

	end := start + req.Size
	if end > total {
		end = total
	}
	kvKeys.Items = keys[start:end]
	return kvKeys, nil
}

func (d *LocalClient) BulkSetValue(ctx context.Context, req *models.BulkSet) (int64, error) {
	var success int64
	for i := range req.Items {
		ok, _ := d.SetValue(ctx, &models.SetValue{
			NamespaceId: req.NamespaceId,
			Key:         req.Items[i].Key,
			Value:       req.Items[i].Value,
			Expiration:  req.Items[i].Expiration,
		})
		if ok {
			success++
		}
	}

	return success, nil
}

func (d *LocalClient) DelValue(ctx context.Context, namespaceId string, key string) (bool, error) {
	file := fmt.Sprintf("%s.json", key)
	if file == metadataFile {
		return true, nil
	}
	path := filepath.Join(storageDir, keyValueDir, namespaceId, file)
	err := os.Remove(path)
	if err != nil {
		return false, fmt.Errorf("delete file %s failed: %v", path, err)
	}

	return true, nil
}

func (d *LocalClient) BulkDelValue(ctx context.Context, namespaceId string, keys []string) (bool, error) {
	for i := range keys {
		_, err := d.DelValue(ctx, namespaceId, keys[i])
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func (d *LocalClient) GetValue(ctx context.Context, namespaceId string, key string) (string, error) {
	file := fmt.Sprintf("%s.json", key)
	path := filepath.Join(storageDir, keyValueDir, namespaceId, file)
	buff, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read file %s failed: %v", path, err)
	}
	var kv models.SetValueLocal
	if err := json.Unmarshal(buff, &kv); err != nil {
		return "", fmt.Errorf("json unmarshal failed: %s", err)
	}
	if kv.ExpireAt.Before(time.Now()) {
		return "", nil
	}
	return kv.Value, nil
}
