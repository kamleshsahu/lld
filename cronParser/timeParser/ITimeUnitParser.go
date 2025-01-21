package timeParser

import "cronParser/entity"

type ITimeUnitParser interface {
	Parse(token string, expression *entity.Expression) error
}

func DefaultTimeUnitMap() []ITimeUnitParser {
	return []ITimeUnitParser{NewMinuteParser(), NewHourParser(), NewDayParser(), NewMonthParser(), NewDayOfWeekParser()}
}
