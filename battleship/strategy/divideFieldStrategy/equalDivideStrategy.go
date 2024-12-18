package divideFieldStrategy

import (
	"lld/battleship/entity"
)

type EqualDivideStrategy struct {
}

func (d *EqualDivideStrategy) Divide(board *entity.Board, players []*entity.Player) error {
	n := len(board.GetCells())
	if n%len(players) != 0 {
		return entity.ErrEqualDivide(n, len(players))
	}
	size := n / len(players)

	for j, row := range board.GetCells() {
		for i, col := range row {
			playerId := i / size
			board.GetCells()[j][i].Owner = players[playerId]
			players[playerId].Field.Cells = append(players[playerId].Field.Cells, *col)
		}
	}
	return nil
}

func NewEqualDivideStrategy() IDivideFieldStrategy {
	return &EqualDivideStrategy{}
}
