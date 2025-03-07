package conditions

import "lld/chess/entity"

type NoBaseCondition struct {
}

func (m NoBaseCondition) IsBaseConditionSatisfied(piece entity.Piece) bool {
	return true
}

func NewNoBaseCondition() entity.IMoveBaseCondition {
	return &NoBaseCondition{}
}
