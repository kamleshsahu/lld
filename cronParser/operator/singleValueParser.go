package operator

import (
	"strconv"
)

type SingleValueParser struct {
}

func (d SingleValueParser) IsApplicable(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}

func (d SingleValueParser) Execute(token string, low, high int, toNumber func(val string) (int, error)) ([]int, error) {
	num, err := toNumber(token)
	if err != nil {
		return nil, err
	}
	return []int{num}, nil
}

func NewSingleValueParser() IOperator {
	return &SingleValueParser{}
}
