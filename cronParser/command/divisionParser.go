package command

import (
	"strconv"
	"strings"
)

type DivisionParser struct {
}

func (d DivisionParser) IsValid(token string) bool {
	return strings.Contains(token, "/")
}

func (d DivisionParser) Execute(token string, low, high int) []int {
	splitToken := strings.Split(token, "/")

	denom, _ := strconv.Atoi(splitToken[1])
	parts := high / denom
	ans := make([]int, 0)
	for i := low; i < parts; i++ {
		ans = append(ans, denom*i)
	}
	return ans
}

func NewDivisionParser() ICommand {
	return &DivisionParser{}
}
