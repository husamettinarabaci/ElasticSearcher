package internal_application_search_service

import (
	"context"

	idsme "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/model/entity"
	idss "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/service"
)

type SearchService struct {
	service idss.SearchService
}

func NewSearchService(s idss.SearchService) SearchService {
	return SearchService{
		service: s,
	}
}

func (c SearchService) GetResultByRequest(ctx context.Context, req idsme.Request) ([]idsme.Result, error) {
	return c.service.GetResultByRequest(ctx, req)
}
