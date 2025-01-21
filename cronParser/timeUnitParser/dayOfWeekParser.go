package timeUnitParser

import (
	"cronParser/entity"
	"strconv"
)

var (
	dayMap = map[string]int{
		"SUN": 0, "MON": 1, "TUE": 2, "WED": 3, "THU": 4, "FRI": 5, "SAT": 6,
	}
)

type dayOfWeekParser struct {
	*BaseParser
}

func NewDayOfWeekParser() ITimeUnitParser {
	return &dayOfWeekParser{
		BaseParser: NewBaseParser(0, 6, entity.DayOfWeek, nil, stringToNumber).(*BaseParser),
	}
}

func stringToNumber(val string) (int, error) {
	dayMapVal, ok := dayMap[val]
	if !ok {
		return strconv.Atoi(val)
	}
	return dayMapVal, nil
}
