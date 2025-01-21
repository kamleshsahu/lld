package command

import (
	"strconv"
)

type NumberParser struct {
	low  int
	high int
}

func (d NumberParser) IsValid(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}

func (d NumberParser) Execute(token string, low, high int) []int {
	num, _ := strconv.Atoi(token)
	return []int{num}
}

func NewNumberParser() ICommand {
	return &NumberParser{}
}
