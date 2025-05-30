package object

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
	Total   int      `json:"total,omitempty"`
}

type ListObjectsResponse struct {
	Objects []ObjectInfo `json:"objects,omitempty"`
	Total   int          `json:"total,omitempty"`
}
type ObjectInfo struct {
	Id        string `json:"id,omitempty"`
	Path      string `json:"path,omitempty"`
	Size      int    `json:"size,omitempty"`
	Filename  string `json:"filename,omitempty"`
	BucketId  string `json:"bucketId,omitempty"`
	ActorId   string `json:"actorId,omitempty"`
	RunId     string `json:"runId,omitempty"`
	FileType  string `json:"fileType,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}
