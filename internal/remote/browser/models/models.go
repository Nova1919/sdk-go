package models

type CreateBrowserRequest struct {
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

type CreateBrowserResponse struct {
	TaskId      string `json:"taskId,omitempty"`
	DevtoolsUrl string `json:"devtoolsUrl,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Code        int64  `json:"code,omitempty"`
	Message     string `json:"message,omitempty"`
}

type ProxyParams struct {
	Url             string `json:"url,omitempty"`
	ChannelId       string `json:"channelId,omitempty"`
	Country         string `json:"country,omitempty"`
	SessionDuration uint64 `json:"sessionDuration,omitempty"`
	SessionId       string `json:"sessionId,omitempty"`
	Gateway         string `json:"gateway,omitempty"`
}
