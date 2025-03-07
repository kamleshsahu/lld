package moves

import "lld/chess/entity"

type CommonMove struct {
	steps                int
	baseCondition        entity.IMoveBaseCondition
	moveFurtherCondition entity.PieceMoveFurtherCondition
	blocker              entity.PieceCellOccupyBlocker
	direction            entity.Direction
}

func (v CommonMove) GetPossibleMoves(board entity.Board, cell *entity.Cell) []*entity.Cell {
	piece := *cell.Piece

	var moves []*entity.Cell

	noOfSteps := 0
	for nextCell, exist := board.GetNextCell(*cell, v.direction); exist && noOfSteps < v.steps; nextCell, exist = board.GetNextCell(*nextCell, v.direction) {
		if !v.baseCondition.IsBaseConditionSatisfied(piece) {
			break
		}
		if v.blocker.IsCellNonOccupiableForPiece(board, *nextCell, piece) {
			break
		}
		moves = append(moves, nextCell)
		if !v.moveFurtherCondition.CanPieceMoveFurtherFromCell(board, *nextCell, piece) {
			break
		}
		noOfSteps++
	}

	return moves
}

func NewCommonMove(steps int, baseCondition entity.IMoveBaseCondition, moveFurtherCondition entity.PieceMoveFurtherCondition, blocker entity.PieceCellOccupyBlocker, direction entity.Direction) entity.IMoveGenerator {
	return &CommonMove{steps: steps, baseCondition: baseCondition, moveFurtherCondition: moveFurtherCondition, blocker: blocker, direction: direction}
}
