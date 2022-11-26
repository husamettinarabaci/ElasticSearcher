package pkg_driving_dto

import (
	"fmt"

	idsme "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/model/entity"
)

type Result struct {
	Id     string  `json:"id"`
	Score  float64 `json:"score"`
	Source string  `json:"source"`
	Err    error   `json:"err"`
}

func (r Result) String() string {
	return fmt.Sprintf("id: %s, score: %f, source: %s, err: %s", r.Id, r.Score, r.Source, r.Err)
}

func NewResult(id string, score float64, source string, err error) Result {
	return Result{
		Id:     id,
		Score:  score,
		Source: source,
		Err:    err,
	}
}

func (r Result) IsEmpty() bool {
	if r.Id == "" && r.Score == 0 && r.Source == "" && r.Err == nil {
		return true
	}
	return false
}

func (r Result) Equals(res Result) bool {
	if r.Id == res.Id && r.Score == res.Score && r.Source == res.Source && r.Err == res.Err {
		return true
	}
	return false
}

func FromResultEntity(res idsme.Result) Result {
	return Result{
		Id:     res.Id.String(),
		Score:  res.Score.Float64(),
		Source: res.Source.String(),
		Err:    res.Err,
	}
}
