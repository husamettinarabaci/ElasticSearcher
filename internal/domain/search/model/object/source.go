package internal_domain_search_model_object

type Source map[string]interface{}

// TODO: implement
func (s Source) String() string {
	return ""
}

func NewSource(source map[string]interface{}) Source {
	return Source(source)
}

func (s Source) IsEmpty() bool {
	return s.String() == ""
}

func (s Source) Equals(source Source) bool {
	return s.String() == source.String()
}
