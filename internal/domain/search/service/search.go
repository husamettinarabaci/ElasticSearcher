package internal_domain_search_service

import (
	"context"

	idsme "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/model/entity"
)

type SearchService struct {
}

func NewSearchService() SearchService {
	return SearchService{}
}

func (c SearchService) GetResultByRequest(ctx context.Context, req idsme.Request) ([]idsme.Result, error) {
	return []idsme.Result{}, nil
}
