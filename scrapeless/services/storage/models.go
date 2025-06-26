package storage

// Dataset

type ItemsResponse struct {
	Items []map[string]any `json:"items,omitempty"`
	Total int              `json:"total,omitempty"`
}

type ListDatasetsResponse struct {
	Items []DatasetInfo `json:"items,omitempty"`
	Total int64         `json:"total,omitempty"`
}

type DatasetInfo struct {
	Id        string   `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	ActorId   string   `json:"actorId,omitempty"`
	RunId     string   `json:"runId,omitempty"`
	Fields    []string `json:"fields,omitempty"` //Fields in the dataset
	CreatedAt string   `json:"createdAt,omitempty"`
	UpdatedAt string   `json:"updatedAt,omitempty"`
}

type Timestamp struct {
	Seconds int64 `json:"seconds,omitempty"`
	Nanos   int32 `json:"nanos,omitempty"`
}

// KV

type KvNamespace struct {
	Items     []KvNamespaceItem `json:"items"`
	Total     int64             `json:"total"`
	Page      int64             `json:"page"`
	PageSize  int64             `json:"pageSize"`
	TotalPage int64             `json:"totalPage"`
}

type KvNamespaceItem struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ActorId   string `json:"actorId"`
	RunId     string `json:"runId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
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
	Expiration  int    `json:"expiration"`
}

type ListKeyInfo struct {
	NamespaceId string `json:"namespaceId"`
	Page        int    `json:"page"`
	Size        int    `json:"size"`
}

type KvKeys struct {
	Items     []map[string]any `json:"items"`
	Total     int64            `json:"total"`
	Page      int64            `json:"page"`
	PageSize  int64            `json:"pageSize"`
	TotalPage int64            `json:"totalPage"`
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

type NamespacesResponse struct {
	Items []KvNamespaceItem `json:"items,omitempty"`
	Total int64             `json:"total,omitempty"`
}

// Object

type Bucket struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
	ActorId     string `json:"actorId,omitempty"`
	RunId       string `json:"runId,omitempty"`
	Size        int    `json:"size,omitempty"`
}
type ListBucketsResponse struct {
	Buckets []Bucket `json:"buckets,omitempty"`
	Total   int64    `json:"total,omitempty"`
}

type ListObjectsResponse struct {
	Objects []ObjectInfo `json:"objects,omitempty"`
	Total   int64        `json:"total,omitempty"`
}
type ObjectInfo struct {
	Id        string `json:"id,omitempty"`
	Path      string `json:"path,omitempty"`
	Size      int    `json:"size,omitempty"`
	Filename  string `json:"filename,omitempty"`
	BucketId  string `json:"bucketId,omitempty"`
	ActorId   string `json:"actorId,omitempty"`
	RunId     string `json:"runId,omitempty"`
	FileType  string `json:"fileType,omitempty"` // The value of FileType is one of json, html, png
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Queue

type ListQueuesResponse struct {
	Items     []Item `json:"items,omitempty"`
	Total     int64  `json:"total"`
	TotalPage int64  `json:"totalPage"`
	Page      int64  `json:"page"`
	PageSize  int64  `json:"pageSize"`
}

type Item struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	TeamId      string `json:"teamId,omitempty"`
	ActorId     string `json:"actorId,omitempty"`
	RunId       string `json:"runId,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}

type CreateQueueReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PushQueue struct {
	Name     string `json:"name"`
	Payload  []byte `json:"payload"`
	Retry    int64  `json:"retry"`
	Timeout  int64  `json:"timeout"`  // timeout-->[60,300] The execution time after the message is pulled, such as 60 seconds; if it exceeds this time, the message will be reset to the pending pull state; Until the retry count is exceeded or the deadline is exceeded
	Deadline int64  `json:"deadline"` // deadline--> [300,86400] The deadline by which messages can be pulled, such as two hours later. Messages that are not pulled after this time will be set as failed
}

type Msg struct {
	ID        string `json:"id"`
	QueueID   string `json:"queueId"`
	Name      string `json:"name"`
	Payload   string `json:"payload"`
	Timeout   int64  `json:"timeout"`
	Deadline  int64  `json:"deadline"`
	Retry     int64  `json:"retry"`
	Retried   int64  `json:"retried"`
	SuccessAt int64  `json:"successAt"`
	FailedAt  int64  `json:"failedAt"`
	Desc      string `json:"desc"`
}

type GetMsgResponse []*Msg
