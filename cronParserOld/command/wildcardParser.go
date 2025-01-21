package command

import (
	"lld/cronParserOld/utils"
	"strings"
)

type WildCardParser struct {
	low  int
	high int
}

func (w WildCardParser) IsValid(token string) bool {
	return strings.EqualFold(token, "*")
}

func (w WildCardParser) Execute(token string, low, high int) []int {
	if low >= 1970 {
		return utils.YearDefaultList
	}
	return utils.GenericDefaultList[:high]
}

func NewWildCardParser() ICommand {
	return &WildCardParser{}
}
