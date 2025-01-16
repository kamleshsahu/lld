package command

import (
	"lld/cronParser/utils"
	"strings"
)

type AstricParser struct {
	low  int
	high int
}

func (d AstricParser) IsValid(token string) bool {
	return strings.EqualFold(token, "*")
}

func (d AstricParser) Execute(token string, low, high int) []int {
	if low >= 1970 {
		return utils.YearDefaultList
	}
	return utils.GenericDefaultList[:high]
}

func NewAstricParser() ICommand {
	return &AstricParser{}
}
