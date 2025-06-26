package storage_memory

import (
	"encoding/json"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage/models"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"os"
	"path/filepath"
	"time"
)

var storageDir string

const (
	datasetDir  = "datasets"
	keyValueDir = "kv_stores"
	queueDir    = "queues_stores"
	objectDir   = "objects_stores"

	metadataFile = "metadata.json"
	inputJson    = "INPUT.json"
	defaultDir   = "default"
)

var defaultLocalClient *LocalClient

type LocalClient struct{}

func Init() {
	cwd, err := os.Getwd()
	if err != nil {
		panic("Unable to get the current working directory：" + err.Error())
	}
	storageDir = filepath.Join(cwd, "storage")
	err = EnsureDir(storageDir)
	if err != nil {
		log.Warnf("warn create storage dir err: %v", err)
	}
	defaultLocalClient = &LocalClient{}
}

func Default() *LocalClient {
	return defaultLocalClient
}

// EnsureDir Ensure that the directory exists (create if it does not exist)
func EnsureDir(storageDir string) error {
	absPath := filepath.Join(storageDir)
	err := createRoot(absPath)
	path, err := createDir(absPath, keyValueDir)
	createMetadata(path, keyValueDir)
	path, err = createDir(absPath, queueDir)
	createMetadata(path, queueDir)
	path, err = createDir(absPath, datasetDir)
	createMetadata(path, datasetDir)
	path, err = createDir(absPath, objectDir)
	createMetadata(path, objectDir)
	createInput(absPath)
	return err
}

func createMetadata(path, category string) {
	metaPath := filepath.Join(path, metadataFile)
	var meta []byte
	var def = "default"
	var now = time.Now().Format(time.RFC3339)
	switch category {
	case datasetDir:
		datasetData := models.Dataset{
			Id:        def,
			Name:      def,
			ActorId:   def,
			RunId:     def,
			CreatedAt: now,
			UpdatedAt: now,
			Stats:     models.DatasetStats{},
		}
		meta, _ = json.MarshalIndent(datasetData, "", "  ")
	case keyValueDir:
		kvData := models.KvNamespaceItem{
			Id:        def,
			Name:      def,
			RunId:     def,
			ActorId:   def,
			CreatedAt: now,
			UpdatedAt: now,
		}
		meta, _ = json.MarshalIndent(kvData, "", "  ")
	case queueDir:
		queue := models.Queue{
			Id:          def,
			Name:        def,
			TeamId:      def,
			ActorId:     def,
			RunId:       def,
			Description: def,
			CreatedAt:   now,
			UpdatedAt:   now,
			Stats:       models.QueueStats{},
		}
		meta, _ = json.MarshalIndent(queue, "", "  ")
	case objectDir:
		bucket := models.Bucket{
			Id:          def,
			Name:        def,
			Description: def,
			CreatedAt:   now,
			UpdatedAt:   now,
			ActorId:     def,
			RunId:       def,
			Size:        0,
		}
		meta, _ = json.MarshalIndent(bucket, "", "  ")
	}
	err := os.WriteFile(metaPath, meta, os.ModePerm)
	if err != nil {
		log.Warnf("warn create metadata.json failed: %v", err)
	}
}

func createInput(absPath string) {
	inputPath := filepath.Join(absPath, keyValueDir, defaultDir, inputJson)
	if !isFileExists(inputPath) {
		create, err := os.Create(inputPath)
		defer create.Close()
		if err != nil {
			log.Warnf("create INPUT.json failed: %v", err)
		}
	}
}

func createRoot(absPath string) error {
	err := os.MkdirAll(absPath, os.ModeDir)
	if err != nil {
		log.Warnf("warn create storage dir err: %v", err)
	}
	return err
}

func createDir(absPath string, dirName string) (string, error) {
	path := filepath.Join(absPath, dirName, defaultDir)
	err := os.MkdirAll(path, os.ModeDir)
	if err != nil {
		log.Warnf("warn create dir err: %v", err)
	}
	return path, err
}

func (c *LocalClient) Close() error {
	return nil
}
