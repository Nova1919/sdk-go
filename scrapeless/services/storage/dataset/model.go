package dataset

type ItemsResponse struct {
	Items []map[string]any `json:"items,omitempty"`
	Total int              `json:"total,omitempty"`
}

type ListDatasetsResponse struct {
	Items []DatasetInfo `json:"items,omitempty"`
	Total int64         `json:"total,omitempty"`
}

type DatasetInfo struct {
	Id         string   `json:"id,omitempty"`
	Name       string   `json:"name,omitempty"`
	ActorId    string   `json:"actorId,omitempty"`
	RunId      string   `json:"runId,omitempty"`
	Fields     []string `json:"fields,omitempty"`
	CreatedAt  string   `json:"createdAt,omitempty"`
	UpdatedAt  string   `json:"updatedAt,omitempty"`
	AccessedAt string   `json:"accessedAt,omitempty"`
}

type Timestamp struct {
	Seconds int64 `json:"seconds,omitempty"`
	Nanos   int32 `json:"nanos,omitempty"`
}
