package operator

import (
	"regexp"
	"strconv"
	"strings"
)

type StepParser struct {
}

func (s StepParser) IsApplicable(token string) bool {
	re1 := regexp.MustCompile(`^(\*|\d+(-\d+)?)/(\d+)$`)
	re2 := regexp.MustCompile(`^(\*|[A-Za-z]+(-[A-Za-z]+)?)/(\d+)$`)
	return re1.MatchString(token) || re2.MatchString(token)
}

func (s StepParser) Execute(token string, low, high int, toNumber func(val string) (int, error)) ([]int, error) {
	splitToken := strings.Split(token, "/")
	start := low
	end := high
	if strings.Contains(splitToken[0], "-") {
		vals := strings.Split(splitToken[0], "-")
		s, err := toNumber(vals[0])
		if err != nil {
			return nil, err
		}
		start = s

		e, err := toNumber(vals[1])
		if err != nil {
			return nil, err
		}
		end = e
	} else if splitToken[0] != "*" {
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

	for i := start; i <= end; i += step {
		ans = append(ans, i)
	}
	return ans, nil
}

func NewStepParser() IOperator {
	return &StepParser{}
}
