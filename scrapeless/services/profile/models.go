package profile

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
	Page     int64   `json:"page"`
	PageSize int64   `json:"pageSize"`
	Name     *string `json:"name"`
}

type ListProfileResponse struct {
	Items     []ProfileInfo `json:"items"`
	Total     int64         `json:"total"`
	Page      int64         `json:"page"`
	PageSize  int64         `json:"pageSize"`
	TotalPage int64         `json:"totalPage"`
}
