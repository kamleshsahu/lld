package operator

import (
	"strings"
)

type CommaParser struct {
}

func (d CommaParser) IsApplicable(token string) bool {
	return strings.Contains(token, ",")
}

func (d CommaParser) Execute(token string, low, high int, toNumber func(val string) (int, error)) ([]int, error) {
	splitToken := strings.Split(token, ",")

	ans := make([]int, len(splitToken))

	for i := 0; i < len(splitToken); i++ {
		val, err := toNumber(splitToken[i])
		if err != nil {
			return nil, err
		}
		ans[i] = val
	}
	return ans, nil
}

func NewCommaParser() IOperator {
	return &CommaParser{}
}
