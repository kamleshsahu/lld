package timeParser

import (
	"cronParser/entity"
)

func NewMinuteParser() ITimeUnitParser {
	return NewBaseParser(0, 59, entity.Minute)
}
