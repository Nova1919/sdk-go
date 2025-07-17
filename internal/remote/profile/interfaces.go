package profile

import (
	"context"
	profile_http "github.com/scrapeless-ai/sdk-go/internal/remote/profile/http"
	"github.com/scrapeless-ai/sdk-go/internal/remote/profile/models"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

type Profile interface {
	Create(ctx context.Context, name string) (profile *models.ProfileInfo, err error)
	Get(ctx context.Context, profileId string) (profile *models.ProfileInfo, err error)
	List(ctx context.Context, req *models.ListProfileRequest) (resp *models.ListProfileResponse, err error)
	Update(ctx context.Context, profileId string, name string) (resp *models.UpdateProfileRequest, err error)
	Delete(ctx context.Context, profileId string) (resp *models.DeleteProfileResponse, err error)
}

var ClientInterface Profile

func NewClient(serverMode, baseUrl string) {
	switch serverMode {
	case "grpc":
		log.Info("grpc...")
	case "dev":
		log.Info("dev...")
	default:
		profile_http.Init(baseUrl)
		ClientInterface = profile_http.Default()
	}
}
