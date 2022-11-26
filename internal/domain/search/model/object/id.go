package internal_domain_search_model_object

type Id string

func (i Id) String() string {
	return string(i)
}

func NewID(id string) Id {
	return Id(id)
}

func (i Id) IsEmpty() bool {
	return i.String() == ""
}

func (i Id) Equals(id Id) bool {
	return i.String() == id.String()
}
