package divideFieldStrategy

import "lld/battleship/entity"

type IDivideFieldStrategy interface {
	Divide(board *entity.Board, players []*entity.Player) error
}
