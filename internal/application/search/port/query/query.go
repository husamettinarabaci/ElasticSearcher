package internal_application_search_port_query

import (
	"context"

	idsme "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/model/entity"
)

type QueryPort interface {
	Search(ctx context.Context, req idsme.Request) ([]idsme.Result, error)
}
