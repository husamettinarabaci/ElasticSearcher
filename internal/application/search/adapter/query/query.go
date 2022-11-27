package internal_application_search_adapter_query

import (
	"context"

	iass "github.com/husamettinarabaci/ElasticSearcher/internal/application/search/service"
	idsme "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/model/entity"
)

type QueryAdapter struct {
	service iass.QueryHandler
}

func NewQueryAdapter(s iass.QueryHandler) QueryAdapter {
	return QueryAdapter{
		service: s,
	}
}

func (q QueryAdapter) Search(ctx context.Context, req idsme.Request) ([]idsme.Result, error) {
	return q.service.Search(ctx, req)
}
