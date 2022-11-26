package internal_application_search_port_cqhandler

import (
	"context"

	idsme "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/model/entity"
)

type QueryHandler interface {
	GetResultByRequest(ctx context.Context, req idsme.Request) ([]idsme.Result, error)
}
