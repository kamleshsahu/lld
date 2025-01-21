package operator

import (
	"strconv"
)

type NumberParser struct {
}

func (d NumberParser) IsApplicable(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}

func (d NumberParser) Execute(token string, low, high int) ([]int, error) {
	num, err := strconv.Atoi(token)
	if err != nil {
		return nil, err
	}
	return []int{num}, nil
}

func NewNumberParser() IOperator {
	return &NumberParser{}
}
