package command

import (
	"strconv"
	"strings"
)

type RangeParser struct {
	low  int
	high int
}

func (d RangeParser) IsValid(token string) bool {
	return strings.Contains(token, "-")
}

func (d RangeParser) Execute(token string, low, high int) []int {
	splitToken := strings.Split(token, "-")

	part1, _ := strconv.Atoi(splitToken[0])
	part2, _ := strconv.Atoi(splitToken[1])

	ans := make([]int, 0)
	for i := part1; i <= part2; i++ {
		ans = append(ans, i)
	}
	return ans
}

func NewRangeParser() ICommand {
	return &RangeParser{}
}
