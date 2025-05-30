package proxy

type GetProxyRequest struct {
	ApiKey          string `json:"apiKey,omitempty"`
	Country         string `json:"country,omitempty"`
	SessionDuration uint64 `json:"sessionDuration,omitempty"`
	SessionId       string `json:"sessionId,omitempty"`
	Gateway         string `json:"gateway,omitempty"`
	TaskId          string `json:"taskId,omitempty"`
}
