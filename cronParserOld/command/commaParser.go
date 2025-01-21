package command

import (
	"strconv"
	"strings"
)

type CommaParser struct {
	low  int
	high int
}

func (d CommaParser) IsValid(token string) bool {
	return strings.Contains(token, ",")
}

func (d CommaParser) Execute(token string, low, high int) []int {
	splitToken := strings.Split(token, ",")

	ans := make([]int, len(splitToken))

	for i := 0; i < len(splitToken); i++ {
		ans[i], _ = strconv.Atoi(splitToken[i])
	}
	return ans
}

func NewCommaParser() ICommand {
	return &CommaParser{}
}
