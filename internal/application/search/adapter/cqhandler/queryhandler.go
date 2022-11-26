package internal_application_search_adapter_cqhandler

import (
	"context"

	iass "github.com/husamettinarabaci/ElasticSearcher/internal/application/search/service"
	idsme "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/model/entity"
)

type QueryHandler struct {
	service iass.SearchService
}

func NewQueryHandler(s iass.SearchService) QueryHandler {
	return QueryHandler{
		service: s,
	}
}

func (q QueryHandler) GetResultByRequest(ctx context.Context, req idsme.Request) ([]idsme.Result, error) {
	return q.service.GetResultByRequest(ctx, req)
}
