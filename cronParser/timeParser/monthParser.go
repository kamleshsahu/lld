package timeParser

import "cronParser/entity"

func NewMonthParser() ITimeUnitParser {
	return NewBaseParser(1, 12, entity.Month)
}
