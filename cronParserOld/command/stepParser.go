package command

import (
	"strconv"
	"strings"
)

type StepParser struct {
}

func (s StepParser) IsValid(token string) bool {
	return strings.Contains(token, "/")
}

func (s StepParser) Execute(token string, low, high int) []int {
	splitToken := strings.Split(token, "/")

	steps, _ := strconv.Atoi(splitToken[1])
	ans := make([]int, 0)

	for i := low; i < high; i += steps {
		ans = append(ans, i)
	}
	return ans
}

func NewStepParser() ICommand {
	return &StepParser{}
}
