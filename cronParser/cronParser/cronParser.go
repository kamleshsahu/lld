package cronParser

import (
	"cronParser/customError"
	"cronParser/entity"
	"cronParser/timeUnitParser"
	"strings"
)

type CronParser struct {
	timeUnitParser []timeUnitParser.ITimeUnitParser
}

func (c *CronParser) Parse(expression string) (*entity.Expression, error) {
	tokens := strings.Split(expression, " ")

	exp := entity.NewExpression()
	tokens, exp.Command = c.pluckCommand(tokens)

	if err := c.IsValidLength(tokens); err != nil {
		return exp, err
	}

	for i, token := range tokens {
		err := c.timeUnitParser[i].Parse(token, exp)
		if err != nil {
			return nil, err
		}
	}

	return exp, nil
}

func (c *CronParser) pluckCommand(tokens []string) ([]string, string) {
	if len(tokens) < 1 {
		return tokens, ""
	}
	last := len(tokens) - 1
	command := tokens[last]
	tokens = tokens[:last]
	return tokens, command
}

func (c *CronParser) IsValidLength(tokens []string) error {
	if len(tokens) != len(c.timeUnitParser) {
		return customError.ErrInvalidNoOfTokens(len(tokens), len(c.timeUnitParser))
	}
	return nil
}

func NewDefaultCronParser(timeUnitParserMap *[]timeUnitParser.ITimeUnitParser) *CronParser {
	cp := CronParser{}
	cp.timeUnitParser = timeUnitParser.DefaultTimeUnitParserMap()
	if timeUnitParserMap != nil {
		cp.timeUnitParser = *timeUnitParserMap
	}
	return &cp
}
