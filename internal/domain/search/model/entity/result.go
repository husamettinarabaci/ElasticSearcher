package internal_domain_search_model_entity

import idsmo "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/model/object"

type Result struct {
	Id     idsmo.Id
	Score  idsmo.Score
	Source idsmo.Source
	Err    error
}

// TODO: implement
func (r Result) String() string {
	return ""
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
