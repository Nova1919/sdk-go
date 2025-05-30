package actor

type IRunActorData struct {
	ActorId    string     `json:"-"`
	Input      any        `json:"input"`
	RunOptions RunOptions `json:"runOptions"`
}

type RunOptions struct {
	CPU     int    `json:"cpu"`
	Memory  int    `json:"memory"`
	Timeout int    `json:"timeout"`
	Version string `json:"version"`
}

type RunInfo struct {
	ActorID     string          `json:"actorId"`
	ActorName   string          `json:"actorName"`
	FinishedAt  string          `json:"finishedAt"`
	Input       map[string]any  `json:"input"`
	InputSchema map[string]any  `json:"inputSchema"`
	Origin      int             `json:"origin"`
	RunID       string          `json:"runId"`
	RunOptions  ResourceOptions `json:"runOptions"`
	SchedulerID string          `json:"schedulerId"`
	StartedAt   string          `json:"startedAt"`
	Status      string          `json:"status"`
	Storage     StorageInfo     `json:"storage"`
	TeamID      string          `json:"teamId"`
}

type ResourceOptions struct {
	CPU          int    `json:"cpu"`
	Memory       int    `json:"memory"`
	ServerMode   int    `json:"serverMode"`
	SurvivalTime int    `json:"survivalTime"`
	Timeout      int    `json:"timeout"`
	Version      string `json:"version"`
}

type StorageInfo struct {
	BucketID      string `json:"bucketId"`
	DatasetID     string `json:"datasetId"`
	KVNamespaceID string `json:"kvNamespaceId"`
	QueueID       string `json:"queueId"`
}

type IPaginationParams struct {
	Page     uint `json:"page"`
	PageSize uint `json:"pageSize"`
	Desc     bool `json:"desc"`
}

type Payload struct {
	ActorID     string          `json:"actorId"`
	ActorName   string          `json:"actorName"`
	FinishedAt  string          `json:"finishedAt"`
	Input       map[string]any  `json:"input"`
	InputSchema map[string]any  `json:"inputSchema"`
	Origin      int             `json:"origin"`
	RunID       string          `json:"runId"`
	RunOptions  ResourceOptions `json:"runOptions"`
	SchedulerID string          `json:"schedulerId"`
	StartedAt   string          `json:"startedAt"`
	Status      string          `json:"status"`
	Storage     StorageInfo     `json:"storage"`
	TeamID      string          `json:"teamId"`
}
type BuildInfo struct {
	ActorID    string `json:"actorId"`
	BuildID    string `json:"buildId"`
	Duration   int    `json:"duration"`
	FinishedAt string `json:"finishedAt"`
	ImageSize  string `json:"imageSize"`
	RepoID     string `json:"repoId"`
	StartedAt  string `json:"startedAt"`
	Status     string `json:"status"`
	TeamID     string `json:"teamId"`
	Version    string `json:"version"`
}
