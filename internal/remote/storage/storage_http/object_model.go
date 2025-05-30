package storage_http

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
