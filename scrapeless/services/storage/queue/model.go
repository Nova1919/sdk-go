package queue

type ListQueuesResponse struct {
	Items  []Item `json:"items,omitempty"`
	Total  int32  `json:"total,omitempty"`
	Limit  int32  `json:"limit,omitempty"`
	Offset int32  `json:"offset,omitempty"`
}
type Item struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	UserId      string `json:"userId,omitempty"`
	TeamId      string `json:"teamId,omitempty"`
	ActorId     string `json:"actorId,omitempty"`
	RunId       string `json:"runId,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}

type Timestamp struct {
	Seconds int64 `json:"seconds,omitempty"`
	Nanos   int32 `json:"nanos,omitempty"`
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
