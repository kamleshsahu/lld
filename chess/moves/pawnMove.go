package moves

import "lld/chess/entity"

type PawnMoves struct {
}

func NewPawnMoves() entity.IMoveGenerator {
	return &PawnMoves{}
}

func (p PawnMoves) GetPossibleMoves(board entity.Board, cell entity.Cell) []*entity.Cell {
	x := cell.X
	y := cell.Y
	piece := cell.Piece

	var moveGenerator []*entity.Cell

	if y+1 < 8 {
		up, _ := board.GetCell(x, y+1)
		if up.Piece == nil {
			moveGenerator = append(moveGenerator, up)
		}
	}

	if x-1 >= 0 && y+1 < 8 {
		upleft, _ := board.GetCell(x-1, y+1)
		if upleft.HasOpponent(piece) {
			moveGenerator = append(moveGenerator, upleft)
		}
	}

	if x+1 < 8 && y+1 < 8 {
		upright, _ := board.GetCell(x+1, y+1)
		if upright.HasOpponent(piece) {
			moveGenerator = append(moveGenerator, upright)
		}
	}

	return moveGenerator
}
