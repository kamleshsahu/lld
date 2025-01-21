package timeParser

import "cronParser/entity"

func NewHourParser() ITimeUnitParser {
	return NewBaseParser(0, 23, entity.Hour)
}
