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

type Piece struct {
	Color         Color
	MoveGenerator []IMoveGenerator
}
