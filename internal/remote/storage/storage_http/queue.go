package storage_http

import (
	"context"
	"fmt"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage/models"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

func (c *Client) CreateQueue(ctx context.Context, req *models.CreateQueueRequest) (*models.CreateQueueResponse, error) {
	handel, ok := queueHandel[createQueue]
	if !ok {
		return nil, fmt.Errorf("not found handle func")
	}
	handel, err := handel.setReq(req).sendRequest(ctx)
	if err != nil {
		return nil, err
	}

	var resp models.CreateQueueResponse
	err = handel.Unmarshal(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (c *Client) GetQueue(ctx context.Context, req *models.GetQueueRequest) (*models.GetQueueResponse, error) {
	handel, ok := queueHandel[getQueue]
	if !ok {
		return nil, fmt.Errorf("not found handle func")
	}
	handel, err := handel.setReq(req).sendRequest(ctx)
	if err != nil {
		return nil, err
	}

	var resp models.GetQueueResponse
	err = handel.Unmarshal(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (c *Client) GetQueues(ctx context.Context, req *models.GetQueuesRequest) (*models.ListQueuesResponse, error) {
	handel, ok := queueHandel[getQueues]
	if !ok {
		return nil, fmt.Errorf("not found handle func")
	}
	handel, err := handel.setReq(req).sendRequest(ctx)
	if err != nil {
		return nil, err
	}

	var resp models.ListQueuesResponse
	err = handel.Unmarshal(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (c *Client) UpdateQueue(ctx context.Context, req *models.UpdateQueueRequest) error {
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

func (c *Client) DelQueue(ctx context.Context, req *models.DelQueueRequest) error {
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

func (c *Client) CreateMsg(ctx context.Context, req *models.CreateMsgRequest) (*models.CreateMsgResponse, error) {
	handel, ok := queueHandel[createMsg]
	if !ok {
		return nil, fmt.Errorf("not found handle func")
	}
	handel, err := handel.setReq(req).sendRequest(ctx)
	if err != nil {
		return nil, err
	}

	var resp models.CreateMsgResponse
	err = handel.Unmarshal(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (c *Client) GetMsg(ctx context.Context, req *models.GetMsgRequest) (*models.GetMsgResponse, error) {
	handel, ok := queueHandel[getMsg]
	if !ok {
		return nil, fmt.Errorf("not found handle func")
	}
	handel, err := handel.setReq(req).sendRequest(ctx)
	if err != nil {
		return nil, err
	}
	var resp models.GetMsgResponse
	err = handel.Unmarshal(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (c *Client) AckMsg(ctx context.Context, req *models.AckMsgRequest) error {
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
