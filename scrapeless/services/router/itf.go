package router

import "io"

type Router interface {
	Request(keyword string, method string, path string, body io.Reader, headers map[string]string) (data []byte, err error)
	Close() error
}
