package storage

type Client struct {
	Dataset
	KV
	Queue
	Object
}

var ClientInterface *Client

func NewClient(env string) {
	ClientInterface = &Client{
		Dataset: nil,
		KV:      nil,
		Queue:   nil,
		Object:  nil,
	}
}
