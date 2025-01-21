package timeParser

import "cronParser/entity"

func NewDayOfWeekParser() ITimeUnitParser {
	return NewBaseParser(1, 5, entity.DayOfWeek)
}
