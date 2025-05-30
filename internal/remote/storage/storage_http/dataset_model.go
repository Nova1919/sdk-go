package storage_http

type ListDatasetsRequest struct {
	ActorId  *string `json:"actorId,omitempty"`
	RunId    *string `json:"runId,omitempty"`
	Page     int64   `json:"page,omitempty"`
	PageSize int64   `json:"pageSize,omitempty"`
	Desc     bool    `json:"desc,omitempty"`
}

type ListDatasetsResponse struct {
	Items     []Dataset `json:"items,omitempty"`
	Total     int64     `json:"total,omitempty"`
	TotalPage int64     `json:"totalPage,omitempty"`
	Page      int64     `json:"page,omitempty"`
	PageSize  int64     `son:"pageSize,omitempty"`
}

type Dataset struct {
	Id         string       `json:"id,omitempty"`
	Name       string       `json:"name,omitempty"`
	ActorId    string       `json:"actorId,omitempty"`
	RunId      string       `json:"runId,omitempty"`
	Fields     []string     `json:"fields,omitempty"`
	CreatedAt  string       `json:"createdAt,omitempty"`
	UpdatedAt  string       `json:"updatedAt,omitempty"`
	AccessedAt string       `json:"accessedAt,omitempty"`
	Stats      DatasetStats `json:"stats,omitempty"`
}

type DatasetStats struct {
	ItemCount      uint64 `json:"itemCount,omitempty"`
	CleanItemCount uint64 `json:"cleanItemCount,omitempty"`
}

type CreateDatasetRequest struct {
	Name    string  `json:"name,omitempty"`
	ActorId *string `json:"actorId,omitempty"`
	RunId   *string `json:"runId,omitempty"`
}

type GetDataset struct {
	DatasetId string `json:"datasetId"`
	Desc      bool   `json:"desc"`
	Page      int    `json:"page"`
	PageSize  int    `json:"pageSize"`
}

type DatasetItem struct {
	Items     []map[string]any `json:"items,omitempty"`
	Total     int              `json:"total"`
	TotalPage int              `json:"totalPage"`
	Page      int              `json:"page"`
	PageSize  int              `json:"pageSize"`
}
