package storage_http

import (
	"context"
	"fmt"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

func (c *Client) CreateQueue(ctx context.Context, req *CreateQueueRequest) (*CreateQueueResponse, error) {
	handel, ok := queueHandel[createQueue]
	if !ok {
		return nil, fmt.Errorf("not found handle func")
	}
	handel, err := handel.setReq(req).sendRequest(ctx)
	if err != nil {
		return nil, err
	}

	var resp CreateQueueResponse
	err = handel.Unmarshal(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (c *Client) GetQueue(ctx context.Context, req *GetQueueRequest) (*GetQueueResponse, error) {
	handel, ok := queueHandel[getQueue]
	if !ok {
		return nil, fmt.Errorf("not found handle func")
	}
	handel, err := handel.setReq(req).sendRequest(ctx)
	if err != nil {
		return nil, err
	}

	var resp GetQueueResponse
	err = handel.Unmarshal(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (c *Client) GetQueues(ctx context.Context, req *GetQueuesRequest) (*ListQueuesResponse, error) {
	handel, ok := queueHandel[getQueues]
	if !ok {
		return nil, fmt.Errorf("not found handle func")
	}
	handel, err := handel.setReq(req).sendRequest(ctx)
	if err != nil {
		return nil, err
	}

	var resp ListQueuesResponse
	err = handel.Unmarshal(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (c *Client) UpdateQueue(ctx context.Context, req *UpdateQueueRequest) error {
	handel, ok := queueHandel[updateQueue]
	if !ok {
		return fmt.Errorf("not found handle func")
	}
	handel, err := handel.setReq(req).sendRequest(ctx)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (c *Client) DelQueue(ctx context.Context, req *DelQueueRequest) error {
	handel, ok := queueHandel[delQueue]
	if !ok {
		return fmt.Errorf("not found handle func")
	}
	handel, err := handel.setReq(req).sendRequest(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) CreateMsg(ctx context.Context, req *CreateMsgRequest) (*CreateMsgResponse, error) {
	handel, ok := queueHandel[createMsg]
	if !ok {
		return nil, fmt.Errorf("not found handle func")
	}
	handel, err := handel.setReq(req).sendRequest(ctx)
	if err != nil {
		return nil, err
	}

	var resp CreateMsgResponse
	err = handel.Unmarshal(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (c *Client) GetMsg(ctx context.Context, req *GetMsgRequest) (*GetMsgResponse, error) {
	handel, ok := queueHandel[getMsg]
	if !ok {
		return nil, fmt.Errorf("not found handle func")
	}
	handel, err := handel.setReq(req).sendRequest(ctx)
	if err != nil {
		return nil, err
	}
	var resp GetMsgResponse
	err = handel.Unmarshal(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (c *Client) AckMsg(ctx context.Context, req *AckMsgRequest) error {
	handel, ok := queueHandel[ackMsg]
	if !ok {
		return fmt.Errorf("not found handle func")
	}
	handel, err := handel.setReq(req).sendRequest(ctx)
	if err != nil {
		return err
	}

	return nil
}
