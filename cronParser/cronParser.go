package main

import (
	"lld/cronParser/entity"
	"lld/cronParser/parser"
	"strings"
)

func CronParser(expression string) (*entity.Expression, string) {
	parserMap := []parser.IParser{parser.NewSecondParser(), parser.NewMinuteParser(), parser.NewHourParser(), parser.NewDayParser(), parser.NewMonthParser(), parser.NewYearParser(), parser.NewDescriptionParser()}
	tokens := strings.Split(expression, " ")

	exp := entity.NewExpression()
	sb := strings.Builder{}
	for i, token := range tokens {
		sb.WriteString(parserMap[i].Parse(token, exp))
		sb.WriteString("\n")
	}

	return exp, sb.String()
}
