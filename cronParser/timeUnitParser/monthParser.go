package timeUnitParser

import (
	"cronParser/customError"
	"cronParser/entity"
)

var (
	monthMap = map[string]int{
		"JAN": 1, "FEB": 2, "MAR": 3, "APR": 4, "MAY": 5, "JUNE": 6, "JULY": 7, "AUG": 8, "SEP": 9, "OCT": 10, "NOV": 11, "DEC": 12,
		"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "11": 11, "12": 12,
	}
)

type monthParser struct {
	*BaseParser
}

func NewMonthParser() ITimeUnitParser {
	return &monthParser{
		BaseParser: NewBaseParser(1, 12, entity.Month, nil, stringToMonthNumber).(*BaseParser),
	}
}

func stringToMonthNumber(val string) (int, error) {
	monthMapVal, ok := monthMap[val]
	if !ok {
		return 0, customError.ErrParsingToken(string(entity.Month), "invalid month")
	}
	return monthMapVal, nil
}
