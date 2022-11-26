package internal_domain_search_model_object

import "fmt"

type Score float64

func (s Score) String() string {
	return fmt.Sprintf("%f", s)
}

func NewScore(score float64) Score {
	return Score(score)
}

func (s Score) IsEmpty() bool {
	return s.String() == ""
}

func (s Score) Equals(score Score) bool {
	return s.String() == score.String()
}

func (s Score) Float64() float64 {
	return float64(s)
}
