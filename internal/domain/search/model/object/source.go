package internal_domain_search_model_object

import "fmt"

type Source map[string]interface{}

func (s Source) CustomString() string {
	return fmt.Sprintf("%v", s)
}

func NewSource(source map[string]interface{}) Source {
	return Source(source)
}

func (s Source) IsEmpty() bool {
	return s.CustomString() == ""
}

func (s Source) Equals(source Source) bool {
	return s.CustomString() == source.CustomString()
}
