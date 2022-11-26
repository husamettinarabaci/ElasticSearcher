package pkg_driving_dto

import (
	"fmt"

	idsme "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/model/entity"
	idsmo "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/model/object"
)

type Request struct {
	UId   string `json:"uid"`
	Index string `json:"index"`
	Query string `json:"query"`
}

func (r Request) String() string {
	return fmt.Sprintf("uid: %s, index: %s, query: %s", r.UId, r.Index, r.Query)
}

func NewRequest(uid string, index string, query string) Request {
	return Request{
		UId:   uid,
		Index: index,
		Query: query,
	}
}

func (r Request) IsEmpty() bool {
	if r.UId == "" && r.Index == "" && r.Query == "" {
		return true
	}
	return false
}

func (r Request) Equals(req Request) bool {
	if r.UId == req.UId && r.Index == req.Index && r.Query == req.Query {
		return true
	}
	return false
}

func (r Request) ToEntity() idsme.Request {
	return idsme.Request{
		UId:   idsmo.NewUId(r.UId),
		Index: idsmo.NewIndex(r.Index),
		Query: idsmo.NewQuery(r.Query),
	}
}
