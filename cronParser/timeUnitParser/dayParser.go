package timeUnitParser

import (
	"cronParser/entity"
)

type dayParser struct {
	*BaseParser
}

func NewDayParser() ITimeUnitParser {
	return &dayParser{
		BaseParser: NewBaseParser(1, 31, entity.Day, nil, nil).(*BaseParser),
	}
}
