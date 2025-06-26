package storage

import (
	"context"

	"github.com/scrapeless-ai/sdk-go/env"
	"github.com/scrapeless-ai/sdk-go/internal/code"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage"
	"github.com/scrapeless-ai/sdk-go/internal/remote/storage/models"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
)

type Vector struct{}

// ListCollections retrieves a list of vector collections with pagination and sorting options.
// Parameters:
//
//	ctx: The context for the request.
//	page: The page number (minimum 1, defaults to 1 if invalid).
//	pageSize: Number of items per page (minimum 10, defaults to 10 if invalid).
//	desc: Whether to sort results in descending order.
func (s *Vector) ListCollections(ctx context.Context, page int64, pageSize int64, desc bool) (*ListCollectionsResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 10 {
		pageSize = 10
	}
	resp, err := storage.ClientInterface.ListCollections(ctx, &models.ListCollectionsRequest{
		Page:     page,
		PageSize: pageSize,
		Desc:     desc,
	})
	if err != nil {
		log.Errorf("failed to list queues: %v", code.Format(err))
		return nil, code.Format(err)
	}
	var items []Collection
	for _, item := range resp.Items {
		items = append(items, Collection{
			Id:          item.Id,
			Name:        item.Name,
			TeamId:      item.TeamId,
			ActorId:     item.ActorId,
			RunId:       item.RunId,
			Description: item.Description,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			Dimension:   item.Dimension,
			Metric:      item.Metric,
			Stats: Stats{
				Count: item.Stats.Count,
				Size:  item.Stats.Size,
			},
		})
	}
	return &ListCollectionsResponse{
		Items:     items,
		Total:     resp.Total,
		TotalPage: resp.TotalPage,
		Page:      resp.Page,
		PageSize:  resp.PageSize,
	}, nil
}

// CreateCollections creates a new vector collection with the provided request parameters.
// Parameters:
//
//	ctx: The context for the request.
//	req: The request object containing collection configuration details.
func (s *Vector) CreateCollections(ctx context.Context, req CreateCollectionRequest) (*CreateCollectionResponse, error) {
	name := req.Name + "-" + env.GetActorEnv().RunId
	resp, err := storage.ClientInterface.CreateCollections(ctx, &models.CreateCollectionRequest{
		ActorId:     env.GetActorEnv().ActorId,
		RunId:       env.GetActorEnv().RunId,
		Name:        name,
		Description: req.Description,
		Dimension:   req.Dimension,
	})
	if err != nil {
		log.Errorf("failed to create queue: %v", code.Format(err))
		return nil, code.Format(err)
	}

	return &CreateCollectionResponse{
		Coll: Collection{
			Id:          resp.Coll.Id,
			Name:        resp.Coll.Name,
			TeamId:      resp.Coll.TeamId,
			ActorId:     resp.Coll.ActorId,
			RunId:       resp.Coll.RunId,
			Description: resp.Coll.Description,
			CreatedAt:   resp.Coll.CreatedAt,
			UpdatedAt:   resp.Coll.UpdatedAt,
			Dimension:   resp.Coll.Dimension,
			Metric:      resp.Coll.Metric,
			Stats: Stats{
				Count: resp.Coll.Stats.Count,
				Size:  resp.Coll.Stats.Size,
			},
		},
	}, nil
}

// UpdateCollection updates the collection information with the provided name and description.
// Parameters:
//
//	ctx: The context for the request.
//	collId: The ID of the collection to update.
//	name: The new name of the collection.
//	description: The new description of the collection.
func (s *Vector) UpdateCollection(ctx context.Context, collId string, name string, description string) error {
	req := &models.UpdateCollectionRequest{
		CollId:      collId,
		Name:        name,
		Description: description,
	}
	err := storage.ClientInterface.UpdateCollection(ctx, req)
	if err != nil {
		log.Errorf("failed to update collection: %v", code.Format(err))
		return code.Format(err)
	}
	return nil
}

// DelCollection deletes the collection by its ID.
// Parameters:
//
//	ctx: The context for the request.
//	collId: The ID of the collection to delete.
func (s *Vector) DelCollection(ctx context.Context, collId string) error {
	err := storage.ClientInterface.DelCollection(ctx, collId)
	if err != nil {
		log.Errorf("failed to delete collection: %v", code.Format(err))
		return code.Format(err)
	}
	return nil
}

// GetCollection retrieves a collection by its ID.
// Parameters:
//
//	ctx: The context for the request.
//	collId: The ID of the collection to retrieve.
func (s *Vector) GetCollection(ctx context.Context, collId string) (*Collection, error) {
	coll, err := storage.ClientInterface.GetCollection(ctx, collId)
	if err != nil {
		log.Errorf("failed to get collection: %v", code.Format(err))
		return nil, code.Format(err)
	}
	return &Collection{
		Id:          coll.Id,
		Name:        coll.Name,
		TeamId:      coll.TeamId,
		ActorId:     coll.ActorId,
		RunId:       coll.RunId,
		Description: coll.Description,
		CreatedAt:   coll.CreatedAt,
		UpdatedAt:   coll.UpdatedAt,
		Dimension:   coll.Dimension,
		Metric:      coll.Metric,
		Stats: Stats{
			Count: coll.Stats.Count,
			Size:  coll.Stats.Size,
		},
	}, nil
}

// CreateDocs inserts new documents into the collection.
// Parameters:
//
//	ctx: The context for the request.
//	collId: The ID of the collection.
//	docs: The documents to insert.
func (s *Vector) CreateDocs(ctx context.Context, collId string, docs []Doc) (*DocOpResponse, error) {
	var modelDocs []models.Doc
	for _, d := range docs {
		modelDocs = append(modelDocs, models.Doc{
			Vector:       d.Vector,
			Content:      d.Content,
			SparseVector: d.SparseVector,
			Score:        d.Score,
		})
	}
	req := &models.CreateDocsRequest{
		CollId: collId,
		Docs:   modelDocs,
	}
	resp, err := storage.ClientInterface.CreateDocs(ctx, req)
	if err != nil {
		log.Errorf("failed to create docs: %v", code.Format(err))
		return nil, code.Format(err)
	}
	var output []DocOpResult
	for _, r := range resp.Output {
		output = append(output, DocOpResult{
			DocOp:   r.DocOp,
			Id:      r.Id,
			Code:    r.Code,
			Message: r.Message,
		})
	}
	return &DocOpResponse{
		Output: output,
	}, nil
}

// UpdateDocs updates existing documents in the collection.
// Parameters:
//
//	ctx: The context for the request.
//	collId: The ID of the collection.
//	docs: The documents to update.
func (s *Vector) UpdateDocs(ctx context.Context, collId string, docs []Doc) (*DocOpResponse, error) {
	var modelDocs []models.Doc
	for _, d := range docs {
		modelDocs = append(modelDocs, models.Doc{
			ID:           d.ID,
			Vector:       d.Vector,
			Content:      d.Content,
			SparseVector: d.SparseVector,
			Score:        d.Score,
		})
	}
	req := &models.UpdateDocsRequest{
		CollId: collId,
		Docs:   modelDocs,
	}
	resp, err := storage.ClientInterface.UpdateDocs(ctx, req)
	if err != nil {
		log.Errorf("failed to update docs: %v", code.Format(err))
		return nil, code.Format(err)
	}
	var output []DocOpResult
	for _, r := range resp.Output {
		output = append(output, DocOpResult{
			DocOp:   r.DocOp,
			Id:      r.Id,
			Code:    r.Code,
			Message: r.Message,
		})
	}
	return &DocOpResponse{
		Output: output,
	}, nil
}

// UpsertDocs inserts or updates documents in the collection.
// Parameters:
//
//	ctx: The context for the request.
//	collId: The ID of the collection.
//	docs: The documents to upsert.
func (s *Vector) UpsertDocs(ctx context.Context, collId string, docs []Doc) (*DocOpResponse, error) {
	var modelDocs []models.Doc
	for _, d := range docs {
		modelDocs = append(modelDocs, models.Doc{
			ID:           d.ID,
			Vector:       d.Vector,
			Content:      d.Content,
			SparseVector: d.SparseVector,
			Score:        d.Score,
		})
	}
	req := &models.UpsertVectorDocsParam{
		CollId: collId,
		Docs:   modelDocs,
	}
	resp, err := storage.ClientInterface.UpsertDocs(ctx, req)
	if err != nil {
		log.Errorf("failed to upsert docs: %v", code.Format(err))
		return nil, code.Format(err)
	}
	var output []DocOpResult
	for _, r := range resp.Output {
		output = append(output, DocOpResult{
			DocOp:   r.DocOp,
			Id:      r.Id,
			Code:    r.Code,
			Message: r.Message,
		})
	}
	return &DocOpResponse{
		Output: output,
	}, nil
}

// DelDocs deletes documents from the collection by their IDs.
// Parameters:
//
//	ctx: The context for the request.
//	collId: The ID of the collection.
//	ids: The IDs of the documents to delete.
func (s *Vector) DelDocs(ctx context.Context, collId string, ids []string) (*DocOpResponse, error) {
	req := &models.DeleteDocsRequest{
		CollId: collId,
		Ids:    ids,
	}
	resp, err := storage.ClientInterface.DelDocs(ctx, req)
	if err != nil {
		log.Errorf("failed to delete docs: %v", code.Format(err))
		return nil, code.Format(err)
	}
	var output []DocOpResult
	for _, r := range resp.Output {
		output = append(output, DocOpResult{
			DocOp:   r.DocOp,
			Id:      r.Id,
			Code:    r.Code,
			Message: r.Message,
		})
	}
	return &DocOpResponse{
		Output: output,
	}, nil
}

// QueryDocs queries documents in the collection by vector.
// Parameters:
//
//	ctx: The context for the request.
//	query: The param of query.
func (s *Vector) QueryDocs(ctx context.Context, collId string, query QueryVectorParam) ([]Doc, error) {
	req := &models.QueryVectorRequest{
		CollId:         collId,
		Vector:         query.Vector,
		SparseVector:   query.SparseVector,
		Topk:           query.Topk,
		IncludeVector:  query.IncludeVector,
		IncludeContent: query.IncludeContent,
	}
	resp, err := storage.ClientInterface.QueryDocs(ctx, req)
	if err != nil {
		log.Errorf("failed to query docs: %v", code.Format(err))
		return nil, code.Format(err)
	}
	var docs []Doc
	for _, d := range resp {
		docs = append(docs, Doc{
			ID:           d.ID,
			Vector:       d.Vector,
			Content:      d.Content,
			SparseVector: d.SparseVector,
			Score:        d.Score,
		})
	}
	return docs, nil
}

// QueryDocsByIds queries documents in the collection by their IDs.
// Parameters:
//
//	ctx: The context for the request.
//	collId: The ID of the collection.
//	ids: The IDs of the documents to query.
func (s *Vector) QueryDocsByIds(ctx context.Context, collId string, ids []string) (map[string]*Doc, error) {
	req := &models.QueryDocsByIdsRequest{
		CollId: collId,
		Ids:    ids,
	}
	resp, err := storage.ClientInterface.QueryDocsByIds(ctx, req)
	if err != nil {
		log.Errorf("failed to query docs by ids: %v", code.Format(err))
		return nil, code.Format(err)
	}
	result := make(map[string]*Doc)
	for k, v := range resp {
		result[k] = &Doc{
			ID:           v.ID,
			Vector:       v.Vector,
			Content:      v.Content,
			SparseVector: v.SparseVector,
			Score:        v.Score,
		}
	}
	return result, nil
}
