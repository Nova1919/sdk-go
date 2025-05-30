package browser

type CreateBrowserRequest struct {
	ApiKey string            `json:"apiKey,omitempty"`
	Input  map[string]string `json:"input,omitempty"`
	Proxy  *ProxyParams      `json:"proxies,omitempty"`
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
