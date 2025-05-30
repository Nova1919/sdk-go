package storage_http

type GetQueueRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type QueueStats struct {
	Pending int64 `json:"pending,omitempty"`
	Running int64 `json:"running,omitempty"`
	Success int64 `json:"success,omitempty"`
	Failed  int64 `json:"failed,omitempty"`
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
