package storage_http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	request2 "github.com/scrapeless-ai/sdk-go/internal/remote/request"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage/models"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"net/http"
)

type IResponse interface {
	IsErr() bool
	Error() string
	GetData() any
}

type HttpHandle[T IResponse] struct {
	Method         string // http.Method
	Url            string
	Req            any
	NeedMarshalReq bool
	FormatURL      func(h *HttpHandle[T]) (string, error)
	respInfo       T // Compatible with other HTTP interfaces with different response structures
}

var queueHandel map[HandleFuncName]*HttpHandle[request2.RespInfo]

type HandleFuncName string

const (
	// 队列
	createQueue = "createQueue"
	getQueue    = "getQueue"
	getQueues   = "getQueues"
	updateQueue = "updateQueue"
	delQueue    = "delQueue"
	// 队列消息
	createMsg = "createMsg"
	getMsg    = "getMsg"
	ackMsg    = "ackMsg"
)

func (c *Client) regisHttpHandleFunc() {

	queueHandel = map[HandleFuncName]*HttpHandle[request2.RespInfo]{
		createQueue: {
			Method:         http.MethodPost,
			Url:            fmt.Sprintf("%s/api/v1/queue", c.BaseUrl),
			NeedMarshalReq: true,
		},
		getQueue: {
			Method:         http.MethodGet,
			Url:            fmt.Sprintf("%s/api/v1/queue", c.BaseUrl),
			NeedMarshalReq: false,
			FormatURL: func(h *HttpHandle[request2.RespInfo]) (string, error) {
				req, ok := h.Req.(*models.GetQueueRequest)
				if !ok {
					return "", errors.New(fmt.Sprintf("type err need GetQueueRequest, but get %T", h.Req))
				}
				return fmt.Sprintf("%s?id=%s&name=%s", h.Url, req.Id, req.Name), nil
			},
		},
		getQueues: {
			Method:         http.MethodGet,
			Url:            fmt.Sprintf("%s/api/v1/queue/queues", c.BaseUrl),
			NeedMarshalReq: true,
			FormatURL: func(h *HttpHandle[request2.RespInfo]) (string, error) {
				req, ok := h.Req.(*models.GetQueuesRequest)
				if !ok {
					return "", errors.New(fmt.Sprintf("type err need GetQueueRequest, but get %T", h.Req))
				}
				return fmt.Sprintf("%s?desc=%t&page=%d&pageSize=%d", h.Url, req.Desc, req.Page, req.PageSize), nil
			},
		},
		updateQueue: {
			Method:         http.MethodPut,
			Url:            fmt.Sprintf("%s/api/v1/queue", c.BaseUrl),
			NeedMarshalReq: true,
			FormatURL: func(h *HttpHandle[request2.RespInfo]) (string, error) {
				req, ok := h.Req.(*models.UpdateQueueRequest)
				if !ok {
					return "", errors.New(fmt.Sprintf("type err need DelQueueRequest, but get %T", h.Req))
				}
				return fmt.Sprintf("%s/%s", h.Url, req.QueueId), nil
			},
		},
		delQueue: {
			Method: http.MethodDelete,
			Url:    fmt.Sprintf("%s/api/v1/queue", c.BaseUrl),
			FormatURL: func(h *HttpHandle[request2.RespInfo]) (string, error) {
				req, ok := h.Req.(*models.DelQueueRequest)
				if !ok {
					return "", errors.New(fmt.Sprintf("type err need DelQueueRequest, but get %T", h.Req))
				}
				return fmt.Sprintf("%s/%s", h.Url, req.QueueId), nil
			},
		},
		createMsg: {
			Method:         http.MethodPost,
			Url:            fmt.Sprintf("%s/api/v1/queue", c.BaseUrl),
			NeedMarshalReq: true,
			FormatURL: func(h *HttpHandle[request2.RespInfo]) (string, error) {
				req, ok := h.Req.(*models.CreateMsgRequest)
				if !ok {
					return "", errors.New(fmt.Sprintf("type err need DelQueueRequest, but get %T", h.Req))
				}
				return fmt.Sprintf("%s/%s/push", h.Url, req.QueueId), nil
			},
		},
		getMsg: {
			Method: http.MethodGet,
			Url:    fmt.Sprintf("%s/api/v1/queue", c.BaseUrl),
			FormatURL: func(h *HttpHandle[request2.RespInfo]) (string, error) {
				req, ok := h.Req.(*models.GetMsgRequest)
				if !ok {
					return "", errors.New(fmt.Sprintf("type err need DelQueueRequest, but get %T", h.Req))
				}
				return fmt.Sprintf("%s/%s/pull?limit=%d", h.Url, req.QueueId, req.Limit), nil
			},
		},
		ackMsg: {
			Method:         http.MethodPost,
			Url:            fmt.Sprintf("%s/api/v1/queue", c.BaseUrl),
			NeedMarshalReq: true,
			FormatURL: func(h *HttpHandle[request2.RespInfo]) (string, error) {
				req, ok := h.Req.(*models.AckMsgRequest)
				if !ok {
					return "", errors.New(fmt.Sprintf("type err need DelQueueRequest, but get %T", h.Req))
				}
				return fmt.Sprintf("%s/%s/ack/%s", h.Url, req.QueueId, req.MsgId), nil
			},
		},
	}
}
func (h *HttpHandle[T]) setReq(req any) *HttpHandle[T] {
	h.Req = req
	return h
}

func (h *HttpHandle[T]) setRespInfo(info T) *HttpHandle[T] {
	h.respInfo = info
	return h
}

func (h *HttpHandle[T]) Unmarshal(resp any) error {
	if h.respInfo.IsErr() {
		return errors.New(h.respInfo.Error())
	}
	marshal, err := json.Marshal(h.respInfo.GetData())
	if err != nil {
		return err
	}
	err = json.Unmarshal(marshal, resp)
	if err != nil {
		return err
	}
	return nil
}

func (h *HttpHandle[T]) sendRequest(ctx context.Context) (*HttpHandle[T], error) {
	defer h.setReq(nil)

	reqBody := ""
	if h.NeedMarshalReq {
		if h.Req == nil {
			return h, errors.New("req is nil")
		}
		reqM, err := json.Marshal(h.Req)
		if err != nil {
			return h, err
		}
		reqBody = string(reqM)
	}

	url := h.Url
	if h.FormatURL != nil {
		u, err := (h.FormatURL)(h)
		if err != nil {
			return h, err
		}
		url = u
	}

	body, err := request2.Request(ctx, request2.ReqInfo{
		Method:  h.Method,
		Url:     url,
		Body:    reqBody,
		Headers: map[string]string{},
	})
	if err != nil {
		log.Errorf("request err:%v", err)
		return h, err
	}
	log.Infof("request body:%s", body)
	var resp T
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return h, err
	}
	h.setRespInfo(resp)
	if resp.IsErr() {
		return h, errors.New(resp.Error())
	}
	return h, nil
}
