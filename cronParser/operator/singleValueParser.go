package operator

import (
	"regexp"
	"strconv"
)

type SingleValueParser struct {
}

func (d SingleValueParser) IsApplicable(token string) bool {
	_, err := strconv.Atoi(token)
	re := regexp.MustCompile(`^([A-Za-z]+)`)
	return err == nil || re.MatchString(token)
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
