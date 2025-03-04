package moves

import "lld/chess/entity"

type Diagonal1Move struct {
}

func (d Diagonal1Move) GetPossibleMoves(board entity.Board, cell entity.Cell) []*entity.Cell {
	x := cell.X
	y := cell.Y
	piece := cell.Piece

	var moveGenerator []*entity.Cell

	mx := min(8-x, 8-y)

	for i := 1; i < mx; i++ {
		tcell, _ := board.GetCell(x+i, y+i)
		if tcell.Piece == nil {
			moveGenerator = append(moveGenerator, tcell)
		} else {
			if tcell.HasOpponent(piece) {
				moveGenerator = append(moveGenerator, tcell)
			}
			break
		}
	}

	mn := min(x, y)

	for i := 1; i < mn; i++ {
		tcell, _ := board.GetCell(x-i, y-i)
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

type Diagonal2Move struct {
}

func (d Diagonal2Move) GetPossibleMoves(board entity.Board, cell entity.Cell) []*entity.Cell {
	x := cell.X
	y := cell.Y
	piece := cell.Piece

	var moveGenerator []*entity.Cell

	mx := min(x, 8-y)

	for i := 1; i < mx; i++ {
		tcell, _ := board.GetCell(x-i, y+i)
		if tcell.Piece == nil {
			moveGenerator = append(moveGenerator, tcell)
		} else {
			if tcell.HasOpponent(piece) {
				moveGenerator = append(moveGenerator, tcell)
			}
			break
		}
	}

	mn := min(8-x, y)

	for i := 1; i < mn; i++ {
		tcell, _ := board.GetCell(x+i, y-i)
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

func NewDiagonal1Move() entity.IMoveGenerator {
	return &Diagonal1Move{}
}

func NewDiagonal2Move() entity.IMoveGenerator {
	return &Diagonal2Move{}
}
