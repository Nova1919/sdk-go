package browser

import "time"

type Actor struct {
	Input           Input  `json:"input"`
	ProxyCountry    string `json:"proxyCountry"`
	ProxyUrl        string `json:"proxyUrl"`
	ChannelId       string `json:"channelId"`
	SessionDuration uint64 `json:"sessionDuration"`
	SessionId       string `json:"sessionId"`
	Gateway         string `json:"gateway"`
	ProfileId       string `json:"profileId"`
	ProfilePersist  bool   `json:"profilePersist"`
}

type ActorOnce struct {
	Input        Input  `json:"input"`
	ProxyCountry string `json:"proxyCountry"`
}

type Input struct {
	SessionTtl string `json:"session_ttl"`
}

type CreateResp struct {
	DevtoolsUrl string `json:"devtoolsUrl"`
	TaskId      string `json:"taskId"`
}

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
