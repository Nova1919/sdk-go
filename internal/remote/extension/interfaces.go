package extension

import (
	"context"
	extension_http "github.com/smash-hq/sdk-go/internal/remote/extension/http"
	"github.com/smash-hq/sdk-go/internal/remote/extension/models"
	"github.com/smash-hq/sdk-go/scrapeless/log"
)

type Extension interface {
	Upload(ctx context.Context, filePath, pluginName string) (extension *models.UploadExtensionResponse, err error)
	Update(ctx context.Context, extensionId, filePath, pluginName string) (success bool, err error)
	Get(ctx context.Context, extensionId string) (extensionDetail *models.ExtensionDetail, err error)
	List(ctx context.Context) (extensionList []models.ExtensionListItem, err error)
	Delete(ctx context.Context, extensionId string) (success bool, err error)
}

var ClientInterface Extension

func NewClient(serverMode, baseUrl string) {
	switch serverMode {
	case "grpc":
		log.Info("grpc...")
	case "dev":
		log.Info("dev...")
	default:
		extension_http.Init(baseUrl)
		ClientInterface = extension_http.Default()
	}
}
