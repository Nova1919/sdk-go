package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/internal/remote/extension/models"
	request2 "github.com/scrapeless-ai/sdk-go/internal/remote/request"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
)

func (c *Client) Upload(ctx context.Context, filePath, pluginName string) (extension *models.UploadExtensionResponse, err error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileName, err := getFileName(filePath)
	if err != nil {
		log.Errorf("get file name err: %v", err)
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Errorf("open file err: %v", err)
		return nil, err
	}
	defer file.Close()

	// Set correct MIME header manually
	partHeader := textproto.MIMEHeader{}
	partHeader.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, fileName))
	partHeader.Set("Content-Type", "application/zip")

	part, err := writer.CreatePart(partHeader)
	if err != nil {
		log.Errorf("create part err: %v", err)
		return nil, err
	}
	io.Copy(part, file)

	writer.WriteField("name", pluginName)
	writer.Close()

	url := fmt.Sprintf("%s/browser/extensions/upload", c.BaseUrl)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		log.Errorf("new request err: %v", err)
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set(env.Env.HTTPHeader, env.GetActorEnv().ApiKey)

	resp, err := c.client.Do(req)
	if err != nil {
		log.Errorf("request error: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	all, _ := io.ReadAll(resp.Body)
	log.Infof("Upload plugin body: %s", string(all))

	if err = json.Unmarshal(all, &extension); err != nil {
		log.Errorf("json unmarshal err: %v", err)
		return nil, err
	}
	return extension, nil
}
func getFileName(filePath string) (string, error) {
	validSuffixes := []string{".zip"}
	fileSuffix := strings.ToLower(filepath.Ext(filePath))

	isValid := false
	for _, suffix := range validSuffixes {
		if fileSuffix == suffix {
			isValid = true
			break
		}
	}

	if !isValid {
		return "", fmt.Errorf("invalid file suffix: %s. Supported suffixes: %s", fileSuffix, strings.Join(validSuffixes, ", "))
	}

	fileName := filepath.Base(filePath)
	return fileName, nil
}
func (c *Client) Update(ctx context.Context, extensionId, filePath, pluginName string) (success bool, err error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	file, err := os.Open(filePath)
	if err != nil {
		log.Errorf("open file err: %v", err)
		return false, err
	}
	defer file.Close()

	partHeader := textproto.MIMEHeader{}
	partHeader.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, pluginName))
	partHeader.Set("Content-Type", "application/zip")

	part, err := writer.CreatePart(partHeader)
	if err != nil {
		log.Errorf("create part err: %v", err)
		return false, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		log.Errorf("copy file err: %v", err)
		return false, err
	}

	writer.WriteField("name", pluginName)
	writer.Close()

	url := fmt.Sprintf("%s/browser/extensions/%s", c.BaseUrl, extensionId)
	request, err := http.NewRequestWithContext(ctx, http.MethodPut, url, body)
	if err != nil {
		log.Errorf("create request err: %v", err)
		return false, err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set(env.Env.HTTPHeader, env.GetActorEnv().ApiKey)

	resp, err := c.client.Do(request)
	if err != nil {
		log.Errorf("request error: %v", err)
		return false, err
	}
	defer resp.Body.Close()

	all, _ := io.ReadAll(resp.Body)
	log.Infof("Update plugin body: %s", string(all))

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return true, nil
	}

	return false, fmt.Errorf("update failed with status %d: %s", resp.StatusCode, string(all))
}

func (c *Client) Get(ctx context.Context, extensionId string) (extensionDetail *models.ExtensionDetail, err error) {
	response, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodGet,
		Url:     fmt.Sprintf("%s/browser/extensions/%s", c.BaseUrl, extensionId),
		Headers: map[string]string{},
	})
	log.Infof("get plugin result:%s", response)
	if err != nil {
		log.Errorf("get plugin result err:%v", err)
		return nil, err
	}
	extensionDetail = new(models.ExtensionDetail)
	if err = json.Unmarshal([]byte(response), extensionDetail); err != nil {
		log.Errorf("json unmarshal err:%v", err)
		return nil, err
	}
	return extensionDetail, nil
}

func (c *Client) List(ctx context.Context) (extensionList []models.ExtensionListItem, err error) {
	response, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodGet,
		Url:     fmt.Sprintf("%s/browser/extensions/list", c.BaseUrl),
		Headers: map[string]string{},
	})
	if err != nil {
		log.Errorf("get plugin list result err:%v", err)
		return nil, err
	}
	if err = json.Unmarshal([]byte(response), &extensionList); err != nil {
		log.Errorf("json unmarshal err:%v", err)
		return nil, err
	}
	return extensionList, nil
}

func (c *Client) Delete(ctx context.Context, extensionId string) (success bool, err error) {
	response, err := request2.Request(ctx, request2.ReqInfo{
		Method:  http.MethodDelete,
		Url:     fmt.Sprintf("%s/browser/extensions/%s", c.BaseUrl, extensionId),
		Headers: map[string]string{},
	})
	log.Infof("delete plugin result:%s", response)
	if err != nil {
		log.Errorf("delete plugin result err:%v", err)
		return false, err
	}
	return true, nil
}
