package divideFieldStrategy

import "lld/battleship/entity"

type EqualDivideStrategy struct {
}

func (d *EqualDivideStrategy) Divide(board *entity.Board, players []*entity.Player) error {
	totalFieldSize := len(board.GetCells())
	size := totalFieldSize / len(players)

	for x := 0; x < totalFieldSize; x++ {
		playerId := x / size
		for y := 0; y < totalFieldSize; y++ {
			board.GetCells()[y][x].Owner = players[playerId]
			players[playerId].Field.Cells = append(players[playerId].Field.Cells, *board.GetCells()[y][x])
		}
	}
	return nil
}

func NewEqualDivideStrategy() IDivideFieldStrategy {
	return &EqualDivideStrategy{}
}
