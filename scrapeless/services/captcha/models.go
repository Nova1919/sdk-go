package captcha

type CaptchaSolverResp struct {
	Token string `json:"token"`
}

type CaptchaSolverReq struct {
	Actor   string    `json:"actor"`
	Input   Input     `json:"input"`
	Proxy   ProxyInfo `json:"proxies"`
	TimeOut int64     `json:"time_out"`
	TaskId  string    `json:"task_id"`
}

type ProxyInfo struct {
	Url             string `json:"url,omitempty"`
	ChannelId       string `json:"channelId,omitempty"`
	Country         string `json:"country,omitempty"`
	SessionDuration uint64 `json:"sessionDuration,omitempty"`
	SessionId       string `json:"sessionId,omitempty"`
	Gateway         string `json:"gateway,omitempty"`
}
type RecaptchaVersion string

const (
	RecaptchaVersionV2 RecaptchaVersion = "v2"
	RecaptchaVersionV3                  = "v3"
)

type Input struct {
	Version           RecaptchaVersion `json:"version"`
	PageURL           string           `json:"pageURL"`
	SiteKey           string           `json:"siteKey"`
	PageAction        string           `json:"pageAction"`
	Invisible         bool             `json:"invisible"`
	EnterprisePayload string           `json:"enterprisePayload"`
}
