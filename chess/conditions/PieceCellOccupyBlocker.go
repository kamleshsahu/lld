package conditions

import "lld/chess/entity"

type PieceCellOccupyBlocker struct {
}

func (p PieceCellOccupyBlocker) IsCellNonOccupiableForPiece(board entity.Board, cell entity.Cell, piece entity.Piece) bool {
	return cell.Piece != nil && cell.Piece.Color == piece.Color
}

func NewPieceCellOccupyBlocker() entity.PieceCellOccupyBlocker {
	return &PieceCellOccupyBlocker{}
}
