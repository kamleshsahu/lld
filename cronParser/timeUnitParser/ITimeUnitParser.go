package timeUnitParser

import "cronParser/entity"

type ITimeUnitParser interface {
	Parse(token string, expression *entity.Expression) error
}

func DefaultTimeUnitParserMap() []ITimeUnitParser {
	return []ITimeUnitParser{NewMinuteParser(), NewHourParser(), NewDayParser(), NewMonthParser(), NewDayOfWeekParser()}
}
