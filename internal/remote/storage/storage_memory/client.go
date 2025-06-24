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
	err := os.MkdirAll(absPath, os.ModeDir)
	if err != nil {
		log.Warnf("warn create storage dir err: %v", err)
	}
	kvPath := filepath.Join(absPath, keyValueDir)
	err = os.MkdirAll(kvPath, os.ModeDir)
	if err != nil {
		log.Warnf("warn create storage dir err: %v", err)
	}
	queuePath := filepath.Join(absPath, queueDir)
	err = os.MkdirAll(queuePath, os.ModeDir)
	if err != nil {
		log.Warnf("warn create storage dir err: %v", err)
	}
	datasetPath := filepath.Join(absPath, datasetDir)
	err = os.MkdirAll(datasetPath, os.ModeDir)
	if err != nil {
		log.Warnf("warn create storage dir err: %v", err)
	}
	objectPath := filepath.Join(absPath, objectDir)
	err = os.MkdirAll(objectPath, os.ModeDir)
	if err != nil {
		log.Warnf("warn create storage dir err: %v", err)
	}
	return err
}

type metadata struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	AccessedAt string   `json:"accessed_at"`
	CreatedAt  string   `json:"created_at"`
	ModifiedAt string   `json:"modified_at"`
	UserId     string   `json:"user_id"`
	ItemCount  int64    `json:"item_count"`
	Fields     []string `json:"fields"`
}
