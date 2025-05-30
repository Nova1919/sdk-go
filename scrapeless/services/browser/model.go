package browser

type Actor struct {
	Input           Input  `json:"input"`
	ProxyCountry    string `json:"proxyCountry"`
	ProxyUrl        string `json:"proxyUrl"`
	ChannelId       string `json:"channelId"`
	SessionDuration uint64 `json:"sessionDuration"`
	SessionId       string `json:"sessionId"`
	Gateway         string `json:"gateway"`
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
