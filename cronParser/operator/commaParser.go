package operator

import (
	"regexp"
	"strings"
)

type CommaParser struct {
	operators []IOperator
}

func (d CommaParser) IsApplicable(token string) bool {
	re1 := regexp.MustCompile(`^(\d+(-\d+)?)(,(\d+(-\d+)?))*$`)
	re2 := regexp.MustCompile(`^([A-Za-z]+(-[A-Za-z]+)?)(,([A-Za-z]+(-[A-Za-z]+)?))*$`)
	return re1.MatchString(token) || re2.MatchString(token)
}

func (d CommaParser) Execute(token string, low, high int, toNumber func(val string) (int, error)) ([]int, error) {
	splitToken := strings.Split(token, ",")

	ans := make([]int, 0)

	for i := 0; i < len(splitToken); i++ {
		for _, operator := range d.operators {
			if operator.IsApplicable(splitToken[i]) {
				vals, err := operator.Execute(splitToken[i], low, high, toNumber)
				if err != nil {
					return nil, err
				}
				ans = append(ans, vals...)
				break
			}
		}
	}
	return ans, nil
}

func NewCommaParser() IOperator {
	operators := []IOperator{NewRangeParser(), NewSingleValueParser()}
	return &CommaParser{operators: operators}
}
