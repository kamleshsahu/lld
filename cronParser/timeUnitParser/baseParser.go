package timeUnitParser

import (
	"cronParser/customError"
	"cronParser/entity"
	"cronParser/operator"
)

type BaseParser struct {
	low       int
	high      int
	timeUnit  entity.TimeUnit
	operators []operator.IOperator
	isValid   func(vals []int) error
}

func (b *BaseParser) Parse(token string, expression *entity.Expression) error {
	for _, operation := range b.operators {
		if operation.IsApplicable(token) {
			vals, err := operation.Execute(token, b.low, b.high)
			if err != nil {
				return customError.ErrParsingToken(b.timeUnit.String(), err)
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

func NewBaseParser(low, high int, timeUnit entity.TimeUnit, inputOps *[]operator.IOperator) ITimeUnitParser {
	ops := operator.DefaultOperatorList()
	if inputOps != nil {
		ops = *inputOps
	}
	return &BaseParser{low: low, high: high, timeUnit: timeUnit, operators: ops}
}
