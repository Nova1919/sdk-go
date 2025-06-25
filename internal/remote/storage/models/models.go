package models

import "time"

// Dataset

type ListDatasetsRequest struct {
	ActorId  *string `json:"actorId,omitempty"`
	RunId    *string `json:"runId,omitempty"`
	Page     int     `json:"page,omitempty"`
	PageSize int     `json:"pageSize,omitempty"`
	Desc     bool    `json:"desc,omitempty"`
}

type ListDatasetsResponse struct {
	Items     []Dataset `json:"items,omitempty"`
	Total     int       `json:"total,omitempty"`
	TotalPage int       `json:"totalPage,omitempty"`
	Page      int       `json:"page,omitempty"`
	PageSize  int       `son:"pageSize,omitempty"`
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
	Size  uint64 `json:"size,omitempty"`
	Count uint64 `json:"count,omitempty"`
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

// Queue

type GetQueueRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type QueueStats struct {
	Pending int `json:"pending,omitempty"`
	Running int `json:"running,omitempty"`
	Success int `json:"success,omitempty"`
	Failed  int `json:"failed,omitempty"`
}

type Queue struct {
	Id          string     `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	UserId      string     `json:"userId,omitempty"`
	TeamId      string     `json:"teamId,omitempty"`
	ActorId     string     `json:"actorId,omitempty"`
	RunId       string     `json:"runId,omitempty"`
	Description string     `json:"description,omitempty"`
	CreatedAt   string     `json:"createdAt,omitempty"`
	UpdatedAt   string     `json:"updatedAt,omitempty"`
	Stats       QueueStats `json:"stats,omitempty"`
}

type GetQueueResponse struct {
	Queue `json:",inline"`
}

type CreateQueueRequest struct {
	ActorId     string `json:"actorId"`
	Name        string `json:"name"`
	RunId       string `json:"runId"`
	Description string `json:"description"`
}

type CreateQueueResponse struct {
	Id string `json:"id"`
}

type GetQueuesRequest struct {
	Desc     bool  `json:"desc"`
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

type ListQueuesResponse struct {
	Items  []*Queue `json:"items"`
	Total  int32    `json:"total"`
	Limit  int32    `json:"limit"`
	Offset int32    `json:"offset"`
}

type UpdateQueueRequest struct {
	QueueId     string `json:"queueId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DelQueueRequest struct {
	QueueId string `json:"queueId"`
}

type CreateMsgRequest struct {
	QueueId  string `json:"queueId"`
	Name     string `json:"name"`
	PayLoad  string `json:"payload"`
	Retry    int64  `json:"retry"`
	Timeout  int64  `json:"timeout"`
	Deadline int64  `json:"deadline"`
}

type CreateMsgResponse struct {
	MsgId string `json:"msgId"`
}

type GetMsgRequest struct {
	QueueId string `json:"queueId"`
	Limit   int32  `json:"limit"`
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

type AckMsgRequest struct {
	QueueId string `json:"queueId"`
	MsgId   string `json:"msgId"`
}

// Bucket

type Bucket struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	ActorId     string `json:"actorId"`
	RunId       string `json:"runId"`
	Size        int    `json:"size"`
}

type Object struct {
	Buckets   []Bucket `json:"buckets"`
	Total     int      `json:"total"`
	TotalPage int      `json:"totalPage"`
	Page      int      `json:"page"`
	PageSize  int      `json:"pageSize"`
}

type CreateBucketRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ActorId     string `json:"actorId,omitempty"`
	RunId       string `json:"runId,omitempty"`
}

type ListObjectsRequest struct {
	BucketId string `json:"bucketId,omitempty"`
	Search   string `json:"search,omitempty"`
	Page     int64  `json:"page,omitempty"`
	PageSize int64  `json:"pageSize,omitempty"`
}

type ObjectList struct {
	Objects   []BucketObject `json:"objects"`
	Total     int            `json:"total"`
	TotalPage int            `json:"totalPage"`
	Page      int            `json:"page"`
	PageSize  int            `json:"pageSize"`
}

type BucketObject struct {
	Id        string `json:"id"`
	Path      string `json:"path"`
	Size      int    `json:"size"`
	Filename  string `json:"filename"`
	BucketId  string `json:"bucketId"`
	ActorId   string `json:"actorId"`
	RunId     string `json:"runId"`
	FileType  string `json:"fileType"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ObjectRequest struct {
	BucketId string `json:"bucketId,omitempty"`
	ObjectId string `json:"objectId,omitempty"`
}

type PutObjectRequest struct {
	BucketId string `json:"bucketId,omitempty"`
	Filename string `json:"filename,omitempty"`
	Data     []byte `json:"data,omitempty"`
	ActorId  string `json:"actorId,omitempty"`
	RunId    string `json:"runId,omitempty"`
}

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

type SetValueFile struct {
	SetValue
	Size     int       `json:"size"`
	ExpireAt time.Time `json:"expireAt"`
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
