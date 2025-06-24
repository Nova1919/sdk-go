package storage_memory

import (
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"os"
	"path/filepath"
)

var storageDir string

const (
	datasetDir  = "datasets"
	keyValueDir = "kv_stores"
	queueDir    = "queues_stores"
	objectDir   = "objects_stores"

	metadataFile = "metadata.json"
)

var fileClient *FileClient

type FileClient struct {
	dir string
}

func Init() {
	cwd, err := os.Getwd()
	if err != nil {
		panic("Unable to get the current working directory：" + err.Error())
	}
	storageDir = filepath.Join(cwd, "storage")
	fileClient = &FileClient{
		dir: storageDir,
	}
	err = fileClient.EnsureDir()
	if err != nil {
		log.Warnf("warn create storage dir err: %v", err)
	}
}

// EnsureDir Ensure that the directory exists (create if it does not exist)
func (fc *FileClient) EnsureDir() error {
	absPath := filepath.Join(fc.dir)
	err := os.MkdirAll(absPath, 0755)
	if err != nil {
		log.Warnf("warn create storage dir err: %v", err)
	}
	kvPath := filepath.Join(absPath, keyValueDir)
	err = os.MkdirAll(kvPath, 0755)
	if err != nil {
		log.Warnf("warn create storage dir err: %v", err)
	}
	queuePath := filepath.Join(absPath, queueDir)
	err = os.MkdirAll(queuePath, 0755)
	if err != nil {
		log.Warnf("warn create storage dir err: %v", err)
	}
	datasetPath := filepath.Join(absPath, datasetDir)
	err = os.MkdirAll(datasetPath, 0755)
	if err != nil {
		log.Warnf("warn create storage dir err: %v", err)
	}
	objectPath := filepath.Join(absPath, objectDir)
	err = os.MkdirAll(objectPath, 0755)
	if err != nil {
		log.Warnf("warn create storage dir err: %v", err)
	}
	return err
}

// CreateDir Create folder (recursive creation)
func (fc *FileClient) CreateDir(relPath string) error {
	absPath := filepath.Join(fc.dir, relPath)
	return os.MkdirAll(absPath, 0755)
}

// DeleteDir Delete the entire folder and its contents
func (fc *FileClient) DeleteDir(relPath string) error {
	absPath := filepath.Join(fc.dir, relPath)
	return os.RemoveAll(absPath)
}

// CreateFile Create file and write content (automatically create directory)
func (fc *FileClient) CreateFile(relPath string, content []byte) error {
	absPath := filepath.Join(fc.dir, relPath)
	dir := filepath.Dir(absPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(absPath, content, 0644)
}

// DeleteFile Delete the specified file
func (fc *FileClient) DeleteFile(relPath string) error {
	absPath := filepath.Join(fc.dir, relPath)
	return os.Remove(absPath)
}

type metadata struct {
	Id         string      `json:"id"`
	Name       interface{} `json:"name"`
	AccessedAt string      `json:"accessed_at"`
	CreatedAt  string      `json:"created_at"`
	ModifiedAt string      `json:"modified_at"`
	UserId     string      `json:"user_id"`
	ItemCount  int64       `json:"item_count"`
	Fields     []string    `json:"fields"`
}
