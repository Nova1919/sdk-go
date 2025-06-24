package storage_memory

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"os"
	"path/filepath"
	"time"
)

type DatasetLocal struct {
	datasetId string
}

func (d *DatasetLocal) ListDatasets(ctx context.Context, page int64, pageSize int64, desc bool) (*storage.ListDatasetsResponse, error) {
	dirPath := filepath.Join(storageDir, datasetDir)

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed read dir: %v", err)
	}
	var s []storage.DatasetInfo
	for _, entry := range entries {
		name := entry.Name()

		file, err := os.ReadFile(filepath.Join(dirPath, name, metadataFile))
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %v", err)
		}
		var meta = &metadata{}
		if err := json.Unmarshal(file, &meta); err != nil {
			return nil, fmt.Errorf("parse JSON %s failed: %v", name, err)
		}
		if entry.IsDir() {
			s = append(s, storage.DatasetInfo{
				Id:     name,
				Name:   name,
				Fields: meta.Fields,
			})
		}
	}
	return &storage.ListDatasetsResponse{
		Items: s,
		Total: int64(len(s)),
	}, nil
}

func (d *DatasetLocal) CreateDataset(ctx context.Context, name string) (id string, datasetName string, err error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return "", "", fmt.Errorf("create dataset failed, cause: %v", err)
	}
	id = newUUID.String()
	path := filepath.Join(storageDir, datasetDir, id)
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return "", "", fmt.Errorf("create dataset failed, cause: %v", err)
	}
	return id, name, nil
}

func (d *DatasetLocal) UpdateDataset(ctx context.Context, datasetId string, name string) (ok bool, datasetName string, err error) {
	_, err = updateMetadata(datasetId, name)
	if err != nil {
		return false, name, fmt.Errorf("dataset update failed, cause: %v", err)
	}
	return true, name, nil
}

func updateMetadata(datasetId string, name string) (*metadata, error) {
	path := filepath.Join(storageDir, datasetDir, datasetId, metadataFile)
	file, err := os.ReadFile(path)
	var meta = &metadata{}
	if err == nil {
		if err := json.Unmarshal(file, &meta); err != nil {
			return nil, fmt.Errorf("parse JSON %s failed: %v", datasetId, err)
		}
	} else {
		err := os.MkdirAll(filepath.Join(storageDir, datasetDir, datasetId), 0755)
		if err != nil {
			return nil, fmt.Errorf("create dataset failed, cause: %v", err)
		}
		meta.Id = datasetId
		meta.UserId = "1"
		meta.CreatedAt = time.Now().Format("2006-01-02T15:04:05.999999999Z07:00")
	}
	meta.Name = name
	meta.ModifiedAt = time.Now().Format("2006-01-02T15:04:05.999999999Z07:00")
	indent, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		log.Warnf("warn json marshal err: %v", err)
	}
	if err := os.WriteFile(path, indent, 0644); err != nil {
		return nil, fmt.Errorf("write file %s failed: %v", path, err)
	}
	return meta, nil
}

func (d *DatasetLocal) DelDataset(ctx context.Context, datasetId string) (bool, error) {
	absPath := filepath.Join(storageDir, datasetDir, datasetId)
	err := os.RemoveAll(absPath)
	if err != nil {
		return false, fmt.Errorf("delete dataset failed, cause: %v", err)
	}
	return true, nil
}

func (d *DatasetLocal) AddItems(ctx context.Context, datasetId string, items []map[string]any) (bool, error) {
	if datasetId == "" {
		datasetId = d.datasetId
	}
	dirPath := filepath.Join(storageDir, datasetDir, datasetId)
	metadataJson, err := updateMetadata(datasetId, "default")
	if err != nil {
		return false, err
	}

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return false, fmt.Errorf("create dir failed: %v", err)
	}

	var fields []string
	for i, item := range items {
		fileName := fmt.Sprintf("%d.json", i)
		filePath := filepath.Join(dirPath, fileName)
		if len(fields) <= 0 {
			for it := range item {
				fields = append(fields, it)
			}
		}
		// Serialize the map to JSON
		data, err := json.MarshalIndent(item, "", "  ")
		if err != nil {
			return false, fmt.Errorf("json marshal failed at index %d: %v", i, err)
		}

		// write file
		if err := os.WriteFile(filePath, data, 0644); err != nil {
			return false, fmt.Errorf("write file %s failed: %v", fileName, err)
		}
	}
	metadataJson.ModifiedAt = time.Now().Format("2006-01-02T15:04:05.999999999Z07:00")
	metadataJson.AccessedAt = time.Now().Format("2006-01-02T15:04:05.999999999Z07:00")
	metadataJson.ItemCount = int64(len(items))
	metadataJson.Fields = fields
	metaFile := filepath.Join(dirPath, metadataFile)
	data, err := json.MarshalIndent(metadataJson, "", "  ")
	if err != nil {
		log.Warnf("warn json marshal err: %v", err)
	}
	if err := os.WriteFile(metaFile, data, 0644); err != nil {
		return false, fmt.Errorf("write file %s failed: %v", metaFile, err)
	}
	return true, nil
}

func (d *DatasetLocal) GetItems(ctx context.Context, datasetId string, page int, pageSize int, desc bool) (*storage.ItemsResponse, error) {
	dirPath := filepath.Join(storageDir, datasetDir, datasetId)
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("read file failed: %v", err)
	}

	var result []map[string]any

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}

		fullPath := filepath.Join(dirPath, entry.Name())

		data, err := os.ReadFile(fullPath)
		if err != nil {
			return nil, fmt.Errorf("read %s failed: %v", entry.Name(), err)
		}

		var item map[string]any
		if err := json.Unmarshal(data, &item); err != nil {
			return nil, fmt.Errorf("parse JSON %s failed: %v", entry.Name(), err)
		}

		result = append(result, item)
	}
	return &storage.ItemsResponse{
		Items: result,
		Total: len(result),
	}, nil
}

func (d *DatasetLocal) Close() error {
	return nil
}
