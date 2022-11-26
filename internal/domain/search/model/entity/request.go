package internal_domain_search_model_entity

import idsmo "github.com/husamettinarabaci/ElasticSearcher/internal/domain/search/model/object"

type Request struct {
	UId   idsmo.UId
	Index idsmo.Index
	Query idsmo.Query
}

// TODO: implement
func (r Request) String() string {
	return ""
}

func NewRequest(uid idsmo.UId, index idsmo.Index, query idsmo.Query) Request {
	return Request{
		UId:   uid,
		Index: index,
		Query: query,
	}
}

func (r Request) IsEmpty() bool {
	return r.UId.IsEmpty()
}

func (r Request) Equals(req Request) bool {
	return r.UId.Equals(req.UId)
}
