package models

import "time"

type ProfileInfo struct {
	ProfileId    string    `json:"profileId"`
	Name         string    `json:"name"`
	Size         int64     `json:"size"`         // profile data size, unit: bytes
	Count        int64     `json:"count"`        // profile usage times
	LastModifyAt time.Time `json:"lastModifyAt"` // last use time
	CreatedAt    time.Time `json:"createdAt"`
}

type ListProfileRequest struct {
	Name     *string `json:"actorId,omitempty"`
	Page     int64   `json:"page,omitempty"`
	PageSize int64   `json:"pageSize,omitempty"`
}

type ListProfileResponse struct {
	Items     []ProfileInfo `json:"docs"`
	Total     int64         `json:"totalDocs"`
	Page      int64         `json:"page"`
	PageSize  int64         `json:"limit"`
	TotalPage int64         `json:"totalPages"`
}

type DeleteProfileResponse struct {
	Success bool `json:"success"`
}

type UpdateProfileRequest DeleteProfileResponse
