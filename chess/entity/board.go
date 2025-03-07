package entity

import "fmt"

type Board struct {
	n     int
	m     int
	cells [8][8]*Cell
}

func (b *Board) GetCell(x, y int) (*Cell, bool) {
	if x < 0 || x >= 8 || y < 0 || y >= 8 {
		return nil, false
	}
	return b.cells[x][y], true
}

func (b *Board) GetNextCell(curr Cell, direction Direction) (*Cell, bool) {
	return b.GetCell(curr.X+direction.X, curr.Y+direction.Y)
}

type Direction struct {
	X int
	Y int
}

var (
	UP        = Direction{0, 1}
	DOWN      = Direction{0, -1}
	LEFT      = Direction{-1, 0}
	RIGHT     = Direction{1, 0}
	UPLEFT    = Direction{-1, 1}
	UPRIGHT   = Direction{1, 1}
	DOWNLEFT  = Direction{-1, -1}
	DOWNRIGHT = Direction{1, -1}
)

var (
	KNIGHTMOVES = []Direction{{1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}}
)

func NewBoard() Board {
	board := Board{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board.cells[i][j] = &Cell{i, j, nil}
		}
	}
	return board
}

func (b *Board) PrintBoard() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			cell := b.cells[j][i]
			if cell.Piece != nil {
				fmt.Print(cell.Piece.String(), " ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}
