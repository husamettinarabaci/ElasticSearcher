package internal_domain_search_model_object

type Query string

func (q Query) String() string {
	return string(q)
}

func NewQuery(query string) Query {
	return Query(query)
}

func (q Query) IsEmpty() bool {
	return q.String() == ""
}

func (q Query) Equals(query Query) bool {
	return q.String() == query.String()
}
