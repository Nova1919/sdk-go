package storage_memory

import (
	"github.com/smash-hq/sdk-go/scrapeless/log"
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
	err := os.MkdirAll(absPath, os.ModeDir)
	if err != nil {
		log.Warnf("warn create storage dir err: %v", err)
	}
	kvPath := filepath.Join(absPath, keyValueDir, defaultDir)
	err = os.MkdirAll(kvPath, os.ModeDir)
	inputPath := filepath.Join(kvPath, inputJson)
	if !isFileExists(inputPath) {
		create, err := os.Create(inputPath)
		defer create.Close()
		if err != nil {
			log.Warnf("create INPUT.json failed: %v", err)
		}
	}
	if err != nil {
		log.Warnf("warn create storage dir err: %v", err)
	}
	queuePath := filepath.Join(absPath, queueDir, defaultDir)
	err = os.MkdirAll(queuePath, os.ModeDir)
	if err != nil {
		log.Warnf("warn create storage dir err: %v", err)
	}
	datasetPath := filepath.Join(absPath, datasetDir, defaultDir)
	err = os.MkdirAll(datasetPath, os.ModeDir)
	if err != nil {
		log.Warnf("warn create storage dir err: %v", err)
	}
	objectPath := filepath.Join(absPath, objectDir, defaultDir)
	err = os.MkdirAll(objectPath, os.ModeDir)
	if err != nil {
		log.Warnf("warn create storage dir err: %v", err)
	}
	return err
}

func (c *LocalClient) Close() error {
	return nil
}
