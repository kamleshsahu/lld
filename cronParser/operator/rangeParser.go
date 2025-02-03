package operator

import (
	"strings"
)

type RangeParser struct {
}

func (d RangeParser) IsApplicable(token string) bool {
	return strings.Contains(token, "-")
}

func (d RangeParser) Execute(token string, low, high int, toNumber func(val string) (int, error)) ([]int, error) {
	splitToken := strings.Split(token, "-")

	start, err := toNumber(splitToken[0])
	if err != nil {
		return nil, err
	}
	end, err := toNumber(splitToken[1])
	if err != nil {
		return nil, err
	}
	ans := make([]int, 0)

	if start < end {
		for i := start; i <= end; i++ {
			ans = append(ans, i)
		}
	} else {
		for i := start; i <= high; i++ {
			ans = append(ans, i)
		}
		for i := low; i <= end; i++ {
			ans = append(ans, i)
		}
	}
	return ans, nil
}

func NewRangeParser() IOperator {
	return &RangeParser{}
}
