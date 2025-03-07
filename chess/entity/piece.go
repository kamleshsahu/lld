package entity

type Color string

const (
	White Color = "White"
	Black Color = "Black"
)

type PieceType string

const (
	Pawn   PieceType = "pawn"
	Knight PieceType = "knight"
	Bishop PieceType = "bishop"
	Rook   PieceType = "rook"
	Queen  PieceType = "queen"
	King   PieceType = "king"
)

func (p Piece) String() string {
	return string(p.PieceType)[0:1]
}

type Piece struct {
	PieceType     PieceType
	Color         Color
	MoveGenerator []IMoveGenerator
	Moves         int
}
