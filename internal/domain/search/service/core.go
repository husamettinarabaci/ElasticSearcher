package internal_domain_search_service

import (
	"context"

	idsme "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/model/entity"
)

type CoreService struct {
}

func NewCoreService() CoreService {
	return CoreService{}
}

func (c CoreService) CheckRequest(ctx context.Context, req idsme.Request) error {
	return nil
}
