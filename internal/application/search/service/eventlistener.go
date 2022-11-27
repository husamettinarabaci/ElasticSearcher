package internal_application_search_service

import (
	"context"

	iasps "github.com/husamettinarabaci/ElasticSearcher/internal/application/search/port/searchengine"
	idsme "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/model/entity"
)

type EventListener struct {
	searchEnginePort iasps.SearchEnginePort
}

func NewEventListener(s iasps.SearchEnginePort) EventListener {
	return EventListener{
		searchEnginePort: s,
	}
}

func (c EventListener) Search(ctx context.Context, req idsme.Request) ([]idsme.Result, error) {
	return c.searchEnginePort.Search(ctx, req)
}
