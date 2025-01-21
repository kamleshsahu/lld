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

func (s StepParser) Execute(token string, low, high int) ([]int, error) {
	splitToken := strings.Split(token, "/")
	start := low
	if splitToken[0] != "*" {
		val, err := strconv.Atoi(splitToken[0])
		if err != nil {
			return nil, err
		}
		start = val
	}

	steps, err := strconv.Atoi(splitToken[1])
	if err != nil {
		return nil, err
	}
	ans := make([]int, 0)

	for i := start; i < high; i += steps {
		ans = append(ans, i)
	}
	return ans, nil
}

func NewStepParser() IOperator {
	return &StepParser{}
}
