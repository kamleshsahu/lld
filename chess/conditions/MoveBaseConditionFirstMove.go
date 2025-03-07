package conditions

import "lld/chess/entity"

type MoveBaseConditionFirstMove struct {
}

func (m MoveBaseConditionFirstMove) IsBaseConditionSatisfied(piece entity.Piece) bool {
	return piece.Moves == 0
}

func NewMoveBaseConditionFirstMove() entity.IMoveBaseCondition {
	return &MoveBaseConditionFirstMove{}
}
