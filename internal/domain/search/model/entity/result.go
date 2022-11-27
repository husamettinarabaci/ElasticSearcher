package internal_domain_search_model_entity

import (
	"fmt"

	idsmo "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/model/object"
)

type Result struct {
	Id     idsmo.Id
	Score  idsmo.Score
	Source idsmo.Source
	Err    error
}

func (r Result) String() string {
	return fmt.Sprintf("id: %s, score: %f, source: %s, err: %s", r.Id.String(), r.Score.Float64(), r.Source.CustomString(), r.Err)
}

func NewResult(id idsmo.Id, score idsmo.Score, source idsmo.Source, err error) Result {
	return Result{
		Id:     id,
		Score:  score,
		Source: source,
		Err:    err,
	}
}

func (r Result) IsEmpty() bool {
	return r.Id.IsEmpty()
}

func (r Result) Equals(res Result) bool {
	return r.Id.Equals(res.Id)
}
