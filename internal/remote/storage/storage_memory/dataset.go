package storage_memory

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage/models"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type DatasetLocal struct {
	datasetId string
}

func (d *DatasetLocal) ListDatasets(ctx context.Context, page int64, pageSize int64, desc bool) (*models.ListDatasetsResponse, error) {
	dirPath := filepath.Join(storageDir, datasetDir)

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir: %v", err)
	}

	var allDatasets []models.Dataset

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

		var meta models.Dataset
		if err := json.Unmarshal(file, &meta); err != nil {
			continue
		}

		allDatasets = append(allDatasets, models.Dataset{
			Id:     name,
			Name:   meta.Name,
			Fields: meta.Fields,
		})
	}

	// sort
	sort.Slice(allDatasets, func(i, j int) bool {
		if desc {
			return allDatasets[i].Name > allDatasets[j].Name
		}
		return allDatasets[i].Name < allDatasets[j].Name
	})

	total := int64(len(allDatasets))

	// page
	start := (page - 1) * pageSize
	if start > total {
		start = total
	}
	end := start + pageSize
	if end > total {
		end = total
	}

	pagedItems := allDatasets[start:end]

	return &models.ListDatasetsResponse{
		Items: pagedItems,
		Total: total,
	}, nil
}

func (d *DatasetLocal) CreateDataset(ctx context.Context, name string) (id string, datasetName string, err error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return "", "", fmt.Errorf("create dataset failed, cause: %v", err)
	}
	id = newUUID.String()
	path := filepath.Join(storageDir, datasetDir, id)
	err = os.MkdirAll(path, os.ModeDir)
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

func updateMetadata(datasetId string, name string) (*models.Dataset, error) {
	path := filepath.Join(storageDir, datasetDir, datasetId, metadataFile)
	file, err := os.ReadFile(path)
	var meta = &models.Dataset{}
	if err == nil {
		if err := json.Unmarshal(file, &meta); err != nil {
			return nil, fmt.Errorf("parse JSON %s failed: %v", datasetId, err)
		}
	} else {
		err := os.MkdirAll(filepath.Join(storageDir, datasetDir, datasetId), os.ModeDir)
		if err != nil {
			return nil, fmt.Errorf("create dataset failed, cause: %v", err)
		}
		meta.Id = datasetId
		meta.CreatedAt = time.Now().Format(time.RFC3339Nano)
	}
	meta.Name = name
	meta.UpdatedAt = time.Now().Format(time.RFC3339Nano)
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

	if err := os.MkdirAll(dirPath, os.ModeDir); err != nil {
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
	metadataJson.UpdatedAt = time.Now().Format(time.RFC3339Nano)
	metadataJson.AccessedAt = time.Now().Format(time.RFC3339Nano)
	metadataJson.Stats = models.DatasetStats{
		Count: uint64(len(items)),
	}
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

func (d *DatasetLocal) GetItems(ctx context.Context, datasetId string, page int, pageSize int, desc bool) (*models.DatasetItem, error) {
	dirPath := filepath.Join(storageDir, datasetDir, datasetId)

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("read dir failed: %v", err)
	}

	var files []os.DirEntry

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		files = append(files, entry)
	}

	// sort
	sort.Slice(files, func(i, j int) bool {
		if desc {
			return files[i].Name() > files[j].Name()
		}
		return files[i].Name() < files[j].Name()
	})

	total := len(files)

	// page
	start := (page - 1) * pageSize
	if start > total {
		start = total
	}
	end := start + pageSize
	if end > total {
		end = total
	}

	pagedFiles := files[start:end]

	var result []map[string]any

	for _, entry := range pagedFiles {
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

	return &models.DatasetItem{
		Items: result,
		Total: total,
	}, nil
}

func (d *DatasetLocal) Close() error {
	return nil
}
