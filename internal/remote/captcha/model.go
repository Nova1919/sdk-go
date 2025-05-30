package captcha

type CreateTaskRequest struct {
	ApiKey  string         `json:"apiKey,omitempty"`
	Actor   string         `json:"actor,omitempty"`
	Input   map[string]any `json:"input,omitempty"`
	Proxy   *ProxyParams   `json:"proxies,omitempty"`
	Timeout int64          `json:"timeout,omitempty"`
}

type ProxyParams struct {
	Url             string `json:"url,omitempty"`
	ChannelId       string `json:"channelId,omitempty"`
	Country         string `json:"country,omitempty"`
	SessionDuration uint64 `json:"sessionDuration,omitempty"`
	SessionId       string `json:"sessionId,omitempty"`
	Gateway         string `json:"gateway,omitempty"`
}

type GetTaskResultRequest struct {
	ApiKey string `json:"apiKey,omitempty"`
	TaskId string `json:"taskId,omitempty"`
}
