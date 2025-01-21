package timeUnitParser

import "cronParser/entity"

func NewDayParser() ITimeUnitParser {
	return NewBaseParser(1, 31, entity.Day, nil)
}
