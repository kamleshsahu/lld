package entity

type Cell struct {
	X     int
	Y     int
	Piece *Piece
}

func (c *Cell) HasOpponent(piece *Piece) bool {
	return c.Piece != nil && c.Piece.Color != piece.Color
}

func (c *Cell) HasSamePieceType(piece *Piece) bool {
	return c.Piece != nil && c.Piece.Color == piece.Color
}

func (c *Cell) Equals(cell *Cell) bool {
	return c.X == cell.X && c.Y == cell.Y
}
