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
	"strconv"
	"strings"
	"time"
)

func (c *LocalClient) ListDatasets(ctx context.Context, req *models.ListDatasetsRequest) (*models.ListDatasetsResponse, error) {
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
		if req.Desc {
			return allDatasets[i].Name > allDatasets[j].Name
		}
		return allDatasets[i].Name < allDatasets[j].Name
	})

	total := int64(len(allDatasets))

	// page
	start := (req.Page - 1) * req.PageSize
	if start > total {
		start = total
	}
	end := start + req.PageSize
	if end > total {
		end = total
	}

	pagedItems := allDatasets[start:end]

	return &models.ListDatasetsResponse{
		Items:     pagedItems,
		Total:     total,
		Page:      req.Page,
		PageSize:  req.PageSize,
		TotalPage: total/req.PageSize + 1,
	}, nil
}

func (c *LocalClient) CreateDataset(ctx context.Context, req *models.CreateDatasetRequest) (*models.Dataset, error) {
	newUUID, err := uuid.NewUUID()
	var rep = &models.Dataset{}
	rep.Name = req.Name
	if err != nil {
		return rep, fmt.Errorf("create dataset failed, cause: %v", err)
	}
	id := newUUID.String()
	path := filepath.Join(storageDir, datasetDir, id)
	err = os.MkdirAll(path, os.ModePerm)
	rep.Id = id
	rep.CreatedAt = time.Now().Format(time.RFC3339Nano)
	rep.UpdatedAt = time.Now().Format(time.RFC3339Nano)
	updateMetadata(id, req.Name)
	if err != nil {
		return rep, fmt.Errorf("create dataset failed, cause: %v", err)
	}
	return rep, nil
}

func (c *LocalClient) UpdateDataset(ctx context.Context, datasetID string, name string) (ok bool, err error) {
	_, err = updateMetadata(datasetID, name)
	if err != nil {
		return false, fmt.Errorf("dataset update failed, cause: %v", err)
	}
	return true, nil
}

func (c *LocalClient) DelDataset(ctx context.Context, datasetID string) (bool, error) {
	absPath := filepath.Join(storageDir, datasetDir, datasetID)
	err := os.RemoveAll(absPath)
	if err != nil {
		return false, fmt.Errorf("delete dataset failed, cause: %v", err)
	}
	return true, nil
}

func (c *LocalClient) GetDataset(ctx context.Context, req *models.GetDataset) (*models.DatasetItem, error) {
	dirPath := filepath.Join(storageDir, datasetDir, req.DatasetId)
	if isDirExists(dirPath) {
		return nil, ErrResourceNotFound
	}
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
		if req.Desc {
			return files[i].Name() > files[j].Name()
		}
		return files[i].Name() < files[j].Name()
	})

	total := len(files)

	// page
	start := (req.Page - 1) * req.PageSize
	if start > total {
		start = total
	}
	end := start + req.PageSize
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
		Items:     result,
		Total:     total,
		Page:      req.Page,
		PageSize:  req.PageSize,
		TotalPage: total/req.PageSize + 1,
	}, nil
}

func (c *LocalClient) AddDatasetItem(ctx context.Context, datasetId string, items []map[string]any) (bool, error) {
	dirPath := filepath.Join(storageDir, datasetDir, datasetId)
	if !isDirExists(dirPath) {
		return false, ErrResourceNotFound
	}
	meta, err := updateMetadata(datasetId, "default")
	if err != nil {
		return false, err
	}

	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return false, fmt.Errorf("create dir failed: %v", err)
	}

	// get number of file
	maxIndex := 0
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return false, fmt.Errorf("read dir failed: %v", err)
	}
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".json") && f.Name() != metadataFile {
			name := strings.TrimSuffix(f.Name(), ".json")
			if len(name) == 8 {
				if idx, err := strconv.Atoi(name); err == nil && idx > maxIndex {
					maxIndex = idx
				}
			}
		}
	}

	var fields []string
	var newSize uint64 = 0

	for i, item := range items {
		index := maxIndex + i + 1
		fileName := fmt.Sprintf("%08d.json", index)
		filePath := filepath.Join(dirPath, fileName)

		if len(fields) == 0 {
			for key := range item {
				fields = append(fields, key)
			}
		}

		data, err := json.MarshalIndent(item, "", "  ")
		if err != nil {
			return false, fmt.Errorf("json marshal failed at index %d: %v", i, err)
		}

		newSize += uint64(len(data))

		if err := os.WriteFile(filePath, data, 0644); err != nil {
			return false, fmt.Errorf("write file %s failed: %v", fileName, err)
		}
	}

	// update metadata
	meta.UpdatedAt = time.Now().Format(time.RFC3339Nano)
	meta.Stats.Count += uint64(len(items))
	meta.Stats.Size += newSize
	meta.Fields = fields

	metaFile := filepath.Join(dirPath, metadataFile)
	data, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		log.Warnf("warn json marshal err: %v", err)
	}
	if err := os.WriteFile(metaFile, data, 0644); err != nil {
		return false, fmt.Errorf("write file %s failed: %v", metaFile, err)
	}
	return true, nil
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
		err := os.MkdirAll(filepath.Join(storageDir, datasetDir, datasetId), os.ModePerm)
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
