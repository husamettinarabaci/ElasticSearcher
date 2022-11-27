package internal_application_search_port_searchengine

import (
	"context"

	idsme "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/model/entity"
)

type SearchEnginePort interface {
	Search(ctx context.Context, req idsme.Request) ([]idsme.Result, error)
}
