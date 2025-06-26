package extension

import (
	"context"
	"github.com/smash-hq/sdk-go/env"
	"github.com/smash-hq/sdk-go/internal/remote/extension"
	"github.com/smash-hq/sdk-go/scrapeless/log"
)

type Extension struct{}

func NewExtension(serverMode string) *Extension {
	log.Info("Internal Extension init")
	extension.NewClient(serverMode, env.Env.ScrapelessBaseApiUrl)
	return &Extension{}
}

func (e *Extension) Upload(ctx context.Context, filePath, pluginName string) (uploadExtension *UploadExtensionResponse, err error) {
	upload, err := extension.ClientInterface.Upload(ctx, filePath, pluginName)
	if err != nil {
		return nil, err
	}
	return &UploadExtensionResponse{
		ExtensionID: upload.ExtensionID,
		Name:        upload.Name,
		CreatedAt:   upload.CreatedAt,
		UpdatedAt:   upload.UpdatedAt,
	}, nil
}

func (e *Extension) Update(ctx context.Context, extensionId, filePath, pluginName string) (success bool, err error) {
	return extension.ClientInterface.Update(ctx, extensionId, filePath, pluginName)
}

func (e *Extension) Get(ctx context.Context, extensionId string) (extensionDetail *ExtensionDetail, err error) {
	detail, err := extension.ClientInterface.Get(ctx, extensionId)
	if err != nil {
		return nil, err
	}
	return &ExtensionDetail{
		ExtensionID:  detail.ExtensionID,
		Name:         detail.Name,
		ManifestName: detail.ManifestName,
		TeamID:       detail.TeamID,
		CreatedAt:    detail.CreatedAt,
		UpdatedAt:    detail.UpdatedAt,
		Version:      detail.Version,
	}, nil

}

func (e *Extension) List(ctx context.Context) (extensionList []ExtensionListItem, err error) {
	list, err := extension.ClientInterface.List(ctx)
	if err != nil {
		return nil, err
	}
	for _, item := range list {
		extensionList = append(extensionList, ExtensionListItem{
			ExtensionID: item.ExtensionID,
			Name:        item.Name,
			Version:     item.Version,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}
	return extensionList, nil
}

func (e *Extension) Delete(ctx context.Context, extensionId string) (success bool, err error) {
	return extension.ClientInterface.Delete(ctx, extensionId)
}

func (e *Extension) Close() error {
	return nil
}
