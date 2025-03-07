package entity

type IMoveGenerator interface {
	GetPossibleMoves(board Board, cell *Cell) []*Cell
}

type IMoveBaseCondition interface {
	IsBaseConditionSatisfied(piece Piece) bool
}

type PieceCellOccupyBlocker interface {
	IsCellNonOccupiableForPiece(Board, Cell, Piece) bool
}

type PieceMoveFurtherCondition interface {
	CanPieceMoveFurtherFromCell(Board, Cell, Piece) bool
}
