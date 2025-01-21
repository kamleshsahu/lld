package timeParser

import (
	"cronParser/entity"
	"cronParser/operator"
	"fmt"
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
				return fmt.Errorf("%s : %s", b.timeUnit, err)
			}
			err = b.isWithinRange(b.low, b.high, vals)
			if err != nil {
				return err
			}
			expression.Set(vals, b.timeUnit)
			return nil
		}
	}

	return fmt.Errorf("%s : %s", b.timeUnit, "no matching opertion")
}

func (b *BaseParser) isWithinRange(low int, high int, arr []int) error {
	if len(arr) == 0 {
		return fmt.Errorf("%s : empty", b.timeUnit)
	}
	for _, val := range arr {
		if val < low || val > high {
			return fmt.Errorf("%s : invalid values", b.timeUnit)
		}
	}
	return nil
}

func NewBaseParser(low, high int, timeUnit entity.TimeUnit) ITimeUnitParser {
	return &BaseParser{low: low, high: high, timeUnit: timeUnit, operators: operator.DefaultOperatorList()}
}
