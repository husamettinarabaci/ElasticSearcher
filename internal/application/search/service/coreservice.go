package internal_application_search_service

import (
	"context"

	idsme "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/model/entity"
	idss "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/service"
)

type CoreService struct {
	domainService idss.CoreService
	eventListener EventListener
}

func NewCoreService(s idss.CoreService, e EventListener) CoreService {
	return CoreService{
		domainService: s,
		eventListener: e,
	}
}

func (c CoreService) Search(ctx context.Context, req idsme.Request) ([]idsme.Result, error) {
	err := c.domainService.CheckRequest(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return c.eventListener.Search(ctx, req)
	}
}
