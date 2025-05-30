package storage_http

type KvNamespace struct {
	Items     []KvNamespaceItem `json:"items"`
	Total     int64             `json:"total"`
	Page      int64             `json:"page"`
	PageSize  int64             `json:"pageSize"`
	TotalPage int64             `json:"totalPage"`
}

type KvNamespaceItem struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	ActorId    string `json:"actorId"`
	RunId      string `json:"runId"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	AccessedAt string `json:"accessedAt"`
}

type CreateKvNamespaceRequest struct {
	Name    string `json:"name"`
	ActorId string `json:"actorId"`
	RunId   string `json:"runId"`
}

type SetValue struct {
	NamespaceId string `json:"namespaceId"`
	Key         string `json:"key"`
	Value       string `json:"value"`
	Expiration  uint   `json:"expiration"`
}

type ListKeyInfo struct {
	NamespaceId string `json:"namespaceId"`
	Page        int    `json:"page"`
	Size        int    `json:"size"`
}

type KvKeys struct {
	Items     []map[string]any `json:"items"`
	Total     int              `json:"total"`
	Page      int              `json:"page"`
	PageSize  int              `json:"pageSize"`
	TotalPage int              `json:"totalPage"`
}

type BulkSet struct {
	NamespaceId string     `json:"namespaceId"`
	Items       []BulkItem `json:"items"`
}

type BulkItem struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Expiration uint   `json:"expiration"`
}
