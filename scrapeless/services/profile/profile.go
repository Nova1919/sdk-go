package profile

import (
	"context"
	"errors"
	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/internal/code"
	"github.com/scrapeless-ai/sdk-go/internal/remote/profile"
	"github.com/scrapeless-ai/sdk-go/internal/remote/profile/models"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

type Profile struct{}

func New() *Profile {
	log.Info("Internal Profile init")
	profile.NewClient("http", env.Env.ScrapelessBaseApiUrl)
	return &Profile{}
}

// CreateProfile creates a new profile.
// Parameters:
//
//	ctx: The request context.
//	name: Name of the profile, default 'untitled'.
func (p *Profile) CreateProfile(ctx context.Context, name string) (*ProfileInfo, error) {
	if name == "" {
		name = "untitled"
	}
	resp, err := profile.ClientInterface.Create(ctx, name)
	if err != nil {
		log.Errorf("create profile err:%v", err)
		return nil, code.Format(err)
	}

	return &ProfileInfo{
		ProfileId:    resp.ProfileId,
		Name:         resp.Name,
		Count:        resp.Count,
		Size:         resp.Size,
		LastModifyAt: resp.LastModifyAt,
		CreatedAt:    resp.CreatedAt,
	}, nil
}

// GetProfile get profile info.
// Parameters:
//
//	ctx: The request context.
//	profileId: Id of the profile.
func (p *Profile) GetProfile(ctx context.Context, profileId string) (*ProfileInfo, error) {
	resp, err := profile.ClientInterface.Get(ctx, profileId)
	if err != nil {
		log.Errorf("get profile err:%v", err)
		return nil, code.Format(err)
	}

	return &ProfileInfo{
		ProfileId:    resp.ProfileId,
		Name:         resp.Name,
		Count:        resp.Count,
		Size:         resp.Size,
		LastModifyAt: resp.LastModifyAt,
		CreatedAt:    resp.CreatedAt,
	}, nil
}

// ListProfiles retrieves a list of profiles with pagination and name search.
// Parameters:
//
//	ctx: The context for the request.
//	req: The request parameters for listing profiles.
func (p *Profile) ListProfiles(ctx context.Context, req *ListProfileRequest) (*ListProfileResponse, error) {
	if req == nil {
		return nil, errors.New("req is nil")
	}
	resp, err := profile.ClientInterface.List(ctx, &models.ListProfileRequest{
		Name:     req.Name,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		log.Errorf("list profile err:%v", err)
		return nil, code.Format(err)
	}
	items := make([]ProfileInfo, 0, len(resp.Items))
	for i := range resp.Items {
		items = append(items, ProfileInfo{
			ProfileId:    resp.Items[i].ProfileId,
			Name:         resp.Items[i].Name,
			Count:        resp.Items[i].Count,
			Size:         resp.Items[i].Size,
			LastModifyAt: resp.Items[i].LastModifyAt,
			CreatedAt:    resp.Items[i].CreatedAt,
		})
	}

	return &ListProfileResponse{
		Items:     items,
		Page:      resp.Page,
		PageSize:  resp.PageSize,
		Total:     resp.Total,
		TotalPage: resp.TotalPage,
	}, nil
}

// UpdateProfile update profile's name
// Parameters:
//
//	ctx: The context for the request.
//	profileId: profile's id.
//	name: profile's name.
func (p *Profile) UpdateProfile(ctx context.Context, profileId string, name string) (bool, error) {
	resp, err := profile.ClientInterface.Update(ctx, profileId, name)
	if err != nil {
		log.Errorf("delete profile err:%v", err)
		return false, code.Format(err)
	}

	return resp.Success, nil
}

// DeleteProfile deletes a profile.
// Parameters:
//
//	ctx: The context for the request.
//	profileId: profile's id.
func (p *Profile) DeleteProfile(ctx context.Context, profileId string) (bool, error) {
	resp, err := profile.ClientInterface.Delete(ctx, profileId)
	if err != nil {
		log.Errorf("delete profile err:%v", err)
		return false, code.Format(err)
	}

	return resp.Success, nil
}

func (p *Profile) Close() error {
	return nil
}
