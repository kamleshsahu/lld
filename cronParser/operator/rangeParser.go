package operator

import (
	"strconv"
	"strings"
)

type RangeParser struct {
}

func (d RangeParser) IsApplicable(token string) bool {
	return strings.Contains(token, "-")
}

func (d RangeParser) Execute(token string, low, high int) ([]int, error) {
	splitToken := strings.Split(token, "-")

	start, err := strconv.Atoi(splitToken[0])
	if err != nil {
		return nil, err
	}
	end, err := strconv.Atoi(splitToken[1])
	if err != nil {
		return nil, err
	}
	ans := make([]int, 0)

	for i := start; i <= end; i++ {
		ans = append(ans, i)
	}
	return ans, nil
}

func NewRangeParser() IOperator {
	return &RangeParser{}
}
