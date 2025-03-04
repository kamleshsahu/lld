package entity

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

func NewBoard() Board {
	board := Board{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board.cells[i][j] = &Cell{i, j, nil}
		}
	}
	return board
}
