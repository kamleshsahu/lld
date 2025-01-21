package timeUnitParser

import "cronParser/entity"

type hourParser struct {
	*BaseParser
}

func NewHourParser() ITimeUnitParser {
	return &hourParser{
		BaseParser: NewBaseParser(0, 23, entity.Hour, nil, nil).(*BaseParser),
	}
}
