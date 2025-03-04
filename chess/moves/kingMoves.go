package moves

import "lld/chess/entity"

type KingMoves struct {
}

func (l KingMoves) GetPossibleMoves(board entity.Board, cell entity.Cell) []*entity.Cell {
	x := cell.X
	y := cell.Y
	piece := cell.Piece

	var moveGenerator []*entity.Cell
	if x > 0 {
		left, _ := board.GetCell(x-1, y)
		if !left.HasSamePieceType(piece) {
			moveGenerator = append(moveGenerator, left)
		}
	}

	if x+1 < 8 {
		right, _ := board.GetCell(x+1, y)
		if !right.HasSamePieceType(piece) {
			moveGenerator = append(moveGenerator, right)
		}
	}

	if y+1 < 8 {
		up, _ := board.GetCell(x, y+1)
		if !up.HasSamePieceType(piece) {
			moveGenerator = append(moveGenerator, up)
		}
	}

	return moveGenerator
}
