package internal_domain_search_model_object

type Index string

func (i Index) String() string {
	return string(i)
}

func NewIndex(index string) Index {
	return Index(index)
}

func (i Index) IsEmpty() bool {
	return i.String() == ""
}

func (i Index) Equals(index Index) bool {
	return i.String() == index.String()
}
