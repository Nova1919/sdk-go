package storage_memory

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/smash-hq/sdk-go/internal/remote/storage/models"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"time"
)

func (q *LocalClient) CreateQueue(ctx context.Context, req *models.CreateQueueRequest) (*models.CreateQueueResponse, error) {
	id := uuid.NewString()
	exists, err := isNameExists(filepath.Join(storageDir, queueDir), req.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("queue %s already exists", req.Name)
	}

	path := filepath.Join(storageDir, queueDir, id)
	err = os.MkdirAll(path, os.ModeDir)
	if err != nil {
		return nil, err
	}
	queue := &models.Queue{
		Id:          id,
		Description: req.Description,
		Name:        req.Name,
		RunId:       req.RunId,
		ActorId:     req.ActorId,
		CreatedAt:   time.Now().Format(time.RFC3339Nano),
	}

	if err = q.updateMetadata(queue); err != nil {
		return nil, fmt.Errorf("update metadata failed, err: %v", err)
	}
	return &models.CreateQueueResponse{
		Id: id,
	}, nil
}

func (q *LocalClient) GetQueue(ctx context.Context, req *models.GetQueueRequest) (*models.GetQueueResponse, error) {
	queuePath := filepath.Join(storageDir, queueDir, req.Id)
	ok := isDirExists(queuePath)
	if !ok {
		return nil, ErrResourceNotFound
	}
	metaDataPath := filepath.Join(queuePath, metadataFile)
	buf, err := os.ReadFile(metaDataPath)
	if err != nil {
		return nil, fmt.Errorf("read file %s failed: %v", metaDataPath, err)
	}

	var queue models.Queue
	if err = json.Unmarshal(buf, &queue); err != nil {
		return nil, fmt.Errorf("json unmarshal failed: %s", err)
	}

	return &models.GetQueueResponse{
		Queue: queue,
	}, nil
}

func (q *LocalClient) GetQueues(ctx context.Context, req *models.GetQueuesRequest) (*models.ListQueuesResponse, error) {
	dirPath := filepath.Join(storageDir, queueDir)

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir: %v", err)
	}

	var allNamespaces []*models.Queue

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		name := entry.Name()
		metaPath := filepath.Join(dirPath, name, metadataFile)

		file, err := os.ReadFile(metaPath)
		if err != nil {
			continue
		}

		var meta models.Queue
		if err = json.Unmarshal(file, &meta); err != nil {
			continue
		}

		allNamespaces = append(allNamespaces, &meta)
	}

	// sort
	sort.Slice(allNamespaces, func(i, j int) bool {
		if req.Desc {
			return allNamespaces[i].CreatedAt > allNamespaces[j].CreatedAt
		}
		return allNamespaces[i].CreatedAt < allNamespaces[j].CreatedAt
	})

	total := int64(len(allNamespaces))

	// page
	start := (req.Page - 1) * req.PageSize
	if start > total {
		start = total
	}
	end := start + req.PageSize
	if end > total {
		end = total
	}

	pagedItems := allNamespaces[start:end]

	return &models.ListQueuesResponse{
		Items:     pagedItems,
		Total:     total,
		Page:      req.Page,
		PageSize:  req.PageSize,
		TotalPage: totalPage(total, req.PageSize),
	}, nil
}

func (q *LocalClient) UpdateQueue(ctx context.Context, req *models.UpdateQueueRequest) error {
	queuePath := filepath.Join(storageDir, queueDir, req.QueueId)
	ok := isDirExists(queuePath)
	if !ok {
		return nil
	}

	metaPath := filepath.Join(queuePath, metadataFile)
	buf, err := os.ReadFile(metaPath)
	if err != nil {
		return fmt.Errorf("read file %s failed: %v", metaPath, err)
	}

	var queue models.Queue
	if err = json.Unmarshal(buf, &queue); err != nil {
		return fmt.Errorf("json unmarshal failed: %s", err)
	}

	queue.Name = req.Name
	queue.Description = req.Description

	return q.updateMetadata(&queue)
}

func (q *LocalClient) DelQueue(ctx context.Context, req *models.DelQueueRequest) error {
	queuePath := filepath.Join(storageDir, queueDir, req.QueueId)
	err := os.RemoveAll(queuePath)
	if err != nil {
		return fmt.Errorf("delete queue failed, cause: %v", err)
	}
	return nil
}

func (q *LocalClient) CreateMsg(ctx context.Context, req *models.CreateMsgRequest) (*models.CreateMsgResponse, error) {
	id := uuid.NewString()
	if req.Deadline < time.Now().Unix()+300 {
		return nil, fmt.Errorf("deadline must after now + 300s")
	}

	msgPath := filepath.Join(storageDir, queueDir, req.QueueId, fmt.Sprintf("%s.json", id))
	msg := models.MsgLocal{
		Msg: models.Msg{
			ID:       id,
			QueueID:  req.QueueId,
			Name:     req.Name,
			Payload:  req.PayLoad,
			Deadline: req.Deadline,
			Retry:    req.Retry,
			Timeout:  req.Timeout,
		},
		UpdateTime: time.Now(),
	}

	marshal, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("json marshal failed: %s", err)
	}

	if err = os.WriteFile(msgPath, marshal, os.ModePerm); err != nil {
		return nil, err
	}
	return &models.CreateMsgResponse{
		MsgId: id,
	}, nil
}

func (l *LocalClient) GetMsg(ctx context.Context, req *models.GetMsgRequest) (*models.GetMsgResponse, error) {
	queuePath := filepath.Join(storageDir, queueDir, req.QueueId)

	msgs := make([]*models.MsgLocal, 0)
	now := time.Now()
	err := filepath.WalkDir(queuePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || d.Name() == metadataFile {
			return nil
		}
		msgPath := filepath.Join(queuePath, d.Name())
		buf, err := os.ReadFile(msgPath)
		if err != nil {
			return fmt.Errorf("read file %s failed: %v", path, err)
		}
		var msg models.MsgLocal
		err = json.Unmarshal(buf, &msg)
		if err != nil {
			return fmt.Errorf("json unmarshal failed: %s", err)
		}
		dead := false
		// msg is dead
		if msg.Deadline < now.Unix() {
			msg.FailedAt = now.Unix()
			dead = true
		}
		// max retry
		if msg.ReenterTime.Before(now) && msg.Retried >= msg.Retry {
			msg.FailedAt = now.Unix()
			dead = true
		}
		if dead {
			msg.UpdateTime = now
			//os.WriteFile(msgPath)
		}
		msgs = append(msgs, &msg)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (q *LocalClient) AckMsg(ctx context.Context, req *models.AckMsgRequest) error {
	//TODO implement me
	panic("implement me")
}

func (q *LocalClient) updateMetadata(queue *models.Queue) error {
	path := filepath.Join(storageDir, queueDir, queue.Id, metadataFile)
	marshal, err := json.Marshal(queue)
	if err != nil {
		return fmt.Errorf("json marshal failed: %s", err)
	}
	return os.WriteFile(path, marshal, os.ModePerm)
}
