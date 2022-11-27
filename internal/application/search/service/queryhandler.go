package internal_application_search_service

import (
	"context"

	idsme "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/model/entity"
)

type QueryHandler struct {
	service CoreService
}

func NewQueryHandler(s CoreService) QueryHandler {
	return QueryHandler{
		service: s,
	}
}

func (c QueryHandler) Search(ctx context.Context, req idsme.Request) ([]idsme.Result, error) {
	return c.service.Search(ctx, req)
}
