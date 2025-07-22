package browser

import "time"

type Actor struct {
	ApiKey      string `json:"apiKey,omitempty"`
	SessionName string `json:"session_name,omitempty"`
	SessionTtl  int    `json:"session_ttl,omitempty"`

	//Whether to enable session recording. When enabled,
	//the entire browser session execution process will be automatically recorded,
	//and after the session is completed, it can be replayed and viewed in the historical session list details.
	SessionRecording bool        `json:"session_recording,omitempty"`
	ProxyCountry     string      `json:"proxy_country"`
	ProxyUrl         string      `json:"proxy_url,omitempty"` //for example: http://user:pass@ip:port
	ExtensionIds     string      `json:"extension_ids,omitempty"`
	Fingerprint      Fingerprint `json:"fingerprint"`
}
type Fingerprint struct {
	UserAgent    string       `json:"user_agent"`
	Platform     string       `json:"platform"`
	Screen       Screen       `json:"screen"`
	Localization Localization `json:"localization"`
}

type Screen struct {
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

type Localization struct {
	BasedOnIP bool     `json:"based_on_ip"`
	Timezone  string   `json:"timezone"`            // eg Asia/Shanghai
	Languages []string `json:"languages,omitempty"` //eg en, en-US
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

type CreateBrowserOptions struct {
	SessionName      string                 `json:"session_name,omitempty"`
	SessionTTL       int                    `json:"session_ttl,omitempty"`
	SessionRecording bool                   `json:"session_recording,omitempty"`
	ProxyCountry     string                 `json:"proxy_country,omitempty"`
	ProxyURL         string                 `json:"proxy_url,omitempty"`
	Fingerprint      map[string]interface{} `json:"fingerprint,omitempty"`
	ExtensionIDs     string                 `json:"extension_ids,omitempty"`
}
