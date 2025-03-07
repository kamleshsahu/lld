package factory

import (
	"lld/chess/conditions"
	"lld/chess/entity"
	"lld/chess/moves"
)

func GetDirections(pieceType entity.PieceType, color entity.Color) []entity.Direction {
	switch pieceType {
	case entity.Pawn:
		switch color {
		case entity.White:
			return []entity.Direction{entity.UP}
		case entity.Black:
			return []entity.Direction{entity.DOWN}
		}
	case entity.Knight:
		return []entity.Direction{{1, 2}, {1, -2}, {2, 1}, {2, -1}, {-1, 2}, {-1, -2}, {-2, 1}, {-2, -1}}
	case entity.Bishop:
		return []entity.Direction{entity.UPLEFT, entity.UPRIGHT, entity.DOWNLEFT, entity.DOWNRIGHT}
	case entity.Rook:
		return []entity.Direction{entity.UP, entity.DOWN, entity.LEFT, entity.RIGHT}
	case entity.Queen:
		return []entity.Direction{entity.UP, entity.DOWN, entity.LEFT, entity.RIGHT, entity.UPLEFT, entity.UPRIGHT, entity.DOWNLEFT, entity.DOWNRIGHT}
	case entity.King:
		return []entity.Direction{entity.UP, entity.DOWN, entity.LEFT, entity.RIGHT, entity.UPLEFT, entity.UPRIGHT, entity.DOWNLEFT, entity.DOWNRIGHT}
	}
	return []entity.Direction{}
}

func GetMoveGenerator(board entity.Board, color entity.Color, pieceType entity.PieceType) []entity.IMoveGenerator {
	var mp []entity.IMoveGenerator
	dirs := GetDirections(pieceType, color)
	switch pieceType {
	case entity.Pawn:
		mp = append(mp, moves.NewCommonMove(2, conditions.NewMoveBaseConditionFirstMove(), conditions.NewPieceMoveFurtherCondition(), conditions.NewPieceCellOccupyBlocker(), dirs[0]))
		mp = append(mp, moves.NewCommonMove(1, conditions.NewNoBaseCondition(), conditions.NewPieceMoveFurtherCondition(), conditions.NewPieceCellOccupyBlocker(), dirs[0]))
	case entity.King:
		for _, dir := range dirs {
			mp = append(mp, moves.NewCommonMove(1, conditions.NewNoBaseCondition(), conditions.NewPieceMoveFurtherCondition(), conditions.NewPieceCellOccupyBlocker(), dir))
		}
	case entity.Bishop:
		for _, dir := range dirs {
			mp = append(mp, moves.NewCommonMove(8, conditions.NewNoBaseCondition(), conditions.NewPieceMoveFurtherCondition(), conditions.NewPieceCellOccupyBlocker(), dir))
		}
	case entity.Rook:
		for _, dir := range dirs {
			mp = append(mp, moves.NewCommonMove(8, conditions.NewNoBaseCondition(), conditions.NewPieceMoveFurtherCondition(), conditions.NewPieceCellOccupyBlocker(), dir))
		}
	case entity.Queen:
		for _, dir := range dirs {
			mp = append(mp, moves.NewCommonMove(8, conditions.NewNoBaseCondition(), conditions.NewPieceMoveFurtherCondition(), conditions.NewPieceCellOccupyBlocker(), dir))
		}
	}
	return mp
}

func GetPieceByType(pieceType entity.PieceType, color entity.Color, board entity.Board) entity.Piece {
	switch pieceType {
	case entity.Pawn:
		return entity.Piece{PieceType: entity.Pawn, Color: color, MoveGenerator: GetMoveGenerator(board, color, entity.Pawn)}
	case entity.Knight:
		return entity.Piece{PieceType: entity.Knight, Color: color, MoveGenerator: GetMoveGenerator(board, color, entity.Knight)}
	case entity.Bishop:
		return entity.Piece{PieceType: entity.Bishop, Color: color, MoveGenerator: GetMoveGenerator(board, color, entity.Bishop)}
	case entity.Rook:
		return entity.Piece{PieceType: entity.Rook, Color: color, MoveGenerator: GetMoveGenerator(board, color, entity.Rook)}
	case entity.Queen:
		return entity.Piece{PieceType: entity.Queen, Color: color, MoveGenerator: GetMoveGenerator(board, color, entity.Queen)}
	case entity.King:
		return entity.Piece{PieceType: entity.King, Color: color, MoveGenerator: GetMoveGenerator(board, color, entity.King)}
	}
	return entity.Piece{}
}
