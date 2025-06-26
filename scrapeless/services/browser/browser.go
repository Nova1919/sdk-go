package browser

import (
	"context"
	"fmt"
	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/internal/code"
	"github.com/scrapeless-ai/sdk-go/internal/remote/browser"
	"github.com/scrapeless-ai/sdk-go/internal/remote/browser/http"
	remote_brwoser "github.com/scrapeless-ai/sdk-go/internal/remote/browser/models"
	"github.com/scrapeless-ai/sdk-go/internal/remote/extension"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/url"
	"strings"
)

type Browser struct {
}

func NewBrowser(serverMode string) *Browser {
	log.Info("browser init")
	browser.NewClient(serverMode, env.Env.ScrapelessBrowserUrl)
	extension.NewClient(serverMode, env.Env.ScrapelessBaseApiUrl)
	return &Browser{}
}
func (b *Browser) Create(ctx context.Context, req Actor) (*CreateResp, error) {
	create, err := browser.ClientInterface.ScrapingBrowserCreate(ctx, &remote_brwoser.CreateBrowserRequest{
		ApiKey: env.GetActorEnv().ApiKey,
		Input: map[string]string{
			"session_ttl": req.Input.SessionTtl,
		},
		Proxy: &remote_brwoser.ProxyParams{
			Url:             req.ProxyUrl,
			ChannelId:       req.ChannelId,
			Country:         strings.ToUpper(req.ProxyCountry),
			SessionDuration: req.SessionDuration,
			SessionId:       req.SessionId,
			Gateway:         req.Gateway,
		},
	})
	if err != nil {
		log.Errorf("scraping browser create err:%v", err)
		return nil, code.Format(err)
	}
	if create != nil {
		return &CreateResp{
			DevtoolsUrl: create.DevtoolsUrl,
			TaskId:      create.TaskId,
		}, nil
	}
	return nil, nil
}

func (b *Browser) CreateOnce(ctx context.Context, req ActorOnce) (*CreateResp, error) {
	u, err := url.Parse(env.Env.ScrapelessBrowserUrl)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "parse url error: %s", err.Error())
	}
	devtoolsUrl := fmt.Sprintf("wss://%s/browser", u.Host)
	value := &url.Values{}
	value.Set("token", env.GetActorEnv().ApiKey)
	value.Set("session_ttl", req.Input.SessionTtl)
	value.Set("proxy_country", strings.ToUpper(req.ProxyCountry))
	return &CreateResp{
		DevtoolsUrl: devtoolsUrl + "?" + value.Encode(),
	}, nil
}

// Upload upload extension
func (b *Browser) Upload(ctx context.Context, filePath, pluginName string) (uploadExtension *UploadExtensionResponse, err error) {
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

// Update update extension
func (b *Browser) Update(ctx context.Context, extensionId, filePath, pluginName string) (success bool, err error) {
	return extension.ClientInterface.Update(ctx, extensionId, filePath, pluginName)
}

// Get get extension detail by extensionId
func (b *Browser) Get(ctx context.Context, extensionId string) (extensionDetail *ExtensionDetail, err error) {
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

// List list extension
func (b *Browser) List(ctx context.Context) (extensionList []ExtensionListItem, err error) {
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

// Delete delete extension by extensionId
func (b *Browser) Delete(ctx context.Context, extensionId string) (success bool, err error) {
	return extension.ClientInterface.Delete(ctx, extensionId)
}

func (b *Browser) Close() error {
	return http.Default().Close()
}
