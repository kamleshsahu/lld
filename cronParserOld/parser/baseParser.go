package parser

import (
	"fmt"
	"lld/cronParserOld/command"
	"lld/cronParserOld/entity"
	"strconv"
	"strings"
)

type BaseParser struct {
	low        int
	high       int
	parserType entity.ParserType
	commands   []command.ICommand
}

func (b *BaseParser) Parse(token string, schedule *entity.Expression) string {
	for _, iCommand := range b.commands {
		if iCommand.IsValid(token) {
			vals := iCommand.Execute(token, b.low, b.high)
			schedule.Set(vals, b.parserType)
			return b.Format(vals)
		}
	}
	return token
}

func (b *BaseParser) Format(vals []int) string {
	sb := strings.Builder{}
	sb.WriteString("Every ")

	v := make([]string, 0)
	for _, val := range vals {
		v = append(v, strconv.Itoa(val))
	}
	sb.WriteString(strings.Join(v, ", "))
	sb.WriteString(fmt.Sprintf(" %s", b.parserType.String()))
	return sb.String()
}

func NewBaseParser(low, high int, parserType entity.ParserType) IParser {
	commands := make([]command.ICommand, 0)
	commands = append(commands, command.NewWildCardParser())
	commands = append(commands, command.NewNumberParser())
	commands = append(commands, command.NewStepParser())
	commands = append(commands, command.NewRangeParser())
	commands = append(commands, command.NewCommaParser())

	return &BaseParser{low: low, high: high, parserType: parserType, commands: commands}
}

func NewSecondParser() IParser {
	return NewBaseParser(0, 60, entity.Second)
}

func NewMinuteParser() IParser {
	return NewBaseParser(0, 60, entity.Minute)
}

func NewHourParser() IParser {
	return NewBaseParser(0, 24, entity.Hour)
}

func NewDayParser() IParser {
	return NewBaseParser(0, 31, entity.Day)
}

func NewMonthParser() IParser {
	return NewBaseParser(0, 12, entity.Month)
}

func NewYearParser() IParser {
	return NewBaseParser(1970, 2099, entity.Year)
}

func NewDescriptionParser() IParser {
	return NewBaseParser(0, 0, entity.Description)
}
