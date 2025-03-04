package moves

import "lld/chess/entity"

type HorizontalMove struct {
}

func (l HorizontalMove) GetPossibleMoves(board entity.Board, cell entity.Cell) []*entity.Cell {
	x := cell.X
	y := cell.Y
	piece := cell.Piece

	moveGenerator := []*entity.Cell{}

	for i := y + 1; i < 8; i++ {
		tcell, _ := board.GetCell(x, i)
		if tcell.Piece == nil {
			moveGenerator = append(moveGenerator, tcell)
		} else {
			if tcell.HasOpponent(piece) {
				moveGenerator = append(moveGenerator, tcell)
			}
			break
		}
	}

	for i := 0; i < y; i++ {
		tcell, _ := board.GetCell(x, i)
		if tcell.Piece == nil {
			moveGenerator = append(moveGenerator, tcell)
		} else {
			if tcell.HasOpponent(piece) {
				moveGenerator = append(moveGenerator, tcell)
			}
			break
		}
	}

	return moveGenerator
}

type VerticalMove struct {
}

func (v VerticalMove) GetPossibleMoves(board entity.Board, cell entity.Cell) []*entity.Cell {
	x := cell.X
	y := cell.Y
	piece := cell.Piece

	moveGenerator := []*entity.Cell{}

	for i := x + 1; i < 8; i++ {
		tcell, _ := board.GetCell(i, y)
		if tcell.Piece == nil {
			moveGenerator = append(moveGenerator, tcell)
		} else {
			if tcell.HasOpponent(piece) {
				moveGenerator = append(moveGenerator, tcell)
			}
			break
		}
	}

	for i := 0; i < y; i++ {
		tcell, _ := board.GetCell(x, i)
		if tcell == nil {
			moveGenerator = append(moveGenerator, tcell)
		} else {
			if tcell.HasOpponent(piece) {
				moveGenerator = append(moveGenerator, tcell)
			}
			break
		}
	}

	return moveGenerator
}

func NewHorizontalMove() entity.IMoveGenerator {
	return &HorizontalMove{}
}

func NewVerticalMove() entity.IMoveGenerator {
	return &VerticalMove{}
}
