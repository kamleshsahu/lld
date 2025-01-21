package timeUnitParser

import (
	"cronParser/entity"
)

type minuteParser struct {
	*BaseParser
}

func NewMinuteParser() ITimeUnitParser {
	return &minuteParser{
		BaseParser: NewBaseParser(0, 59, entity.Minute, nil, nil).(*BaseParser),
	}
}
