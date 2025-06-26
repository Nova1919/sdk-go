package models

import "time"

type UploadExtensionResponse struct {
	ExtensionID string    `json:"extensionId"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ExtensionDetail struct {
	ExtensionID  string    `json:"extensionId"`
	TeamID       string    `json:"teamId"`
	ManifestName string    `json:"manifestName"`
	Name         string    `json:"name"`
	Version      string    `json:"version"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type ExtensionListItem struct {
	ExtensionID string    `json:"extensionId"`
	Name        string    `json:"name"`
	Version     string    `json:"version"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
