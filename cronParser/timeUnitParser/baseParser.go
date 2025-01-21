package timeUnitParser

import (
	"cronParser/customError"
	"cronParser/entity"
	"cronParser/operator"
	"strconv"
)

type BaseParser struct {
	low       int
	high      int
	timeUnit  entity.TimeUnit
	operators []operator.IOperator
	parseFn   func(string) (int, error)
}

func (b *BaseParser) Parse(token string, expression *entity.Expression) error {
	for _, operation := range b.operators {
		if operation.IsApplicable(token) {
			vals, err := operation.Execute(token, b.low, b.high, b.StringToNumber)
			if err != nil {
				return customError.ErrParsingToken(b.timeUnit.String(), err.Error())
			}
			err = b.isWithinRange(b.low, b.high, vals)
			if err != nil {
				return err
			}
			expression.Set(vals, b.timeUnit)
			return nil
		}
	}

	return customError.ErrNoMatchingOperation(b.timeUnit.String())
}

func (b *BaseParser) isWithinRange(low int, high int, arr []int) error {
	if len(arr) == 0 {
		return customError.ErrEmptyNumberRange(b.timeUnit.String())
	}
	for _, val := range arr {
		if val < low || val > high {
			return customError.ErrInvalidNumberRange(b.timeUnit.String())
		}
	}
	return nil
}

func (b *BaseParser) StringToNumber(value string) (int, error) {
	if b.parseFn != nil {
		return b.parseFn(value)
	}
	return strconv.Atoi(value)
}

func NewBaseParser(low, high int, timeUnit entity.TimeUnit, inputOps *[]operator.IOperator, parseFn func(string) (int, error)) ITimeUnitParser {
	ops := operator.DefaultOperatorList()
	if inputOps != nil {
		ops = *inputOps
	}
	return &BaseParser{low: low, high: high, timeUnit: timeUnit, operators: ops, parseFn: parseFn}
}
