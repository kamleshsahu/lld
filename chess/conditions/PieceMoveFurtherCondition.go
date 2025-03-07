package conditions

import "lld/chess/entity"

type PieceMoveFurtherCondition struct {
}

func (p PieceMoveFurtherCondition) CanPieceMoveFurtherFromCell(board entity.Board, cell entity.Cell, piece entity.Piece) bool {
	return cell.Piece == nil
}

func NewPieceMoveFurtherCondition() entity.PieceMoveFurtherCondition {
	return &PieceMoveFurtherCondition{}
}
