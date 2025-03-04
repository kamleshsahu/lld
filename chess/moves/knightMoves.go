package moves

import "lld/chess/entity"

type KnightMoves struct {
}

func (p KnightMoves) GetPossibleMoves(board entity.Board, cell entity.Cell) []*entity.Cell {
	x := cell.X
	y := cell.Y
	piece := cell.Piece

	dirs := [][]int{{1, 2}, {1, -2}, {-1, 2}, {-1, -2}, {2, 1}, {2, -1}, {-2, 1}, {-2, -1}}
	var moveGenerator []*entity.Cell
	for _, dir := range dirs {
		newX := x + dir[0]
		newY := y + dir[1]
		if newX >= 0 && newX < 8 && newY >= 0 && newY < 8 {
			tcell, exists := board.GetCell(newX, newY)
			if !exists {
				continue
			}
			if tcell.Piece == nil || tcell.HasOpponent(piece) {
				moveGenerator = append(moveGenerator, tcell)
			}
		}
	}
	return moveGenerator
}

func NewKnightMoves() entity.IMoveGenerator {
	return &KnightMoves{}
}
