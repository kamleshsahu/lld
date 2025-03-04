package entity

type IMoveGenerator interface {
	GetPossibleMoves(board Board, cell Cell) []*Cell
}
