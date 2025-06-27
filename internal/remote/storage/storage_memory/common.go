package storage_memory

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var (
	ErrResourceNotFound          = errors.New("resource not found")
	ErrResourceExists            = errors.New("resource exists")
	ErrLocalStorageUnimplemented = errors.New("local storage unimplemented")
)

func isDirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func isFileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func totalPage(total, pageSize int64) int64 {
	return (total + pageSize - 1) / pageSize
}

func isNameExists(path string, name string) (bool, error) {
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			return nil
		}
		if d.Name() == queueDir || d.Name() == datasetDir || d.Name() == keyValueDir ||
			d.Name() == objectDir || d.Name() == metadataFile {
			return nil
		}
		metaDataPath := filepath.Join(path, metadataFile)
		file, err := os.ReadFile(metaDataPath)
		if err != nil {
			return err
		}
		metaData := make(map[string]any)
		err = json.Unmarshal(file, &metaData)
		if err != nil {
			return fmt.Errorf("json unmarshal failed: %s", err)
		}
		if metaData["name"] == name {
			return ErrResourceExists
		}

		return nil
	})

	if err != nil {
		if errors.Is(err, ErrResourceExists) {
			return true, nil
		}
		return false, err
	}

	return false, nil
}
