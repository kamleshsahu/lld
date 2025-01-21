package operator

import (
	"strconv"
	"strings"
)

type StepParser struct {
}

func (s StepParser) IsApplicable(token string) bool {
	return strings.Contains(token, "/")
}

func (s StepParser) Execute(token string, low, high int, toNumber func(val string) (int, error)) ([]int, error) {
	splitToken := strings.Split(token, "/")
	start := low
	if splitToken[0] != "*" {
		val, err := toNumber(splitToken[0])
		if err != nil {
			return nil, err
		}
		start = val
	}

	step, err := strconv.Atoi(splitToken[1])
	if err != nil {
		return nil, err
	}
	ans := make([]int, 0)

	for i := start; i <= high; i += step {
		ans = append(ans, i)
	}
	return ans, nil
}

func NewStepParser() IOperator {
	return &StepParser{}
}
