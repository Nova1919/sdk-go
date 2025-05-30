package helper

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

// CompressData zpi
func GzipCompressData(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	_, err := writer.Write(data)
	if err != nil {
		return nil, err
	}
	writer.Close()
	return buf.Bytes(), nil
}

// GzipDecompressData unzip
func GzipDecompressData(compressedData []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(compressedData))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	decompressedData, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return decompressedData, nil
}
