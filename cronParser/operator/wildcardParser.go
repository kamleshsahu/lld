package operator

import (
	"strings"
)

type WildCardParser struct {
}

func (w WildCardParser) IsApplicable(token string) bool {
	return strings.EqualFold(token, "*")
}

func (w WildCardParser) Execute(token string, low, high int) ([]int, error) {
	arr := make([]int, 0)

	for i := low; i <= high; i++ {
		arr = append(arr, i)
	}

	return arr, nil
}

func NewWildCardParser() IOperator {
	return &WildCardParser{}
}
