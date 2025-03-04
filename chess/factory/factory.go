package factory

import (
	"lld/chess/entity"
	"lld/chess/moves"
)

func GetMoveGenerator(pieceType entity.PieceType) []entity.IMoveGenerator {
	switch pieceType {
	case entity.Pawn:
		return []entity.IMoveGenerator{moves.NewPawnMoves()}
	case entity.Knight:
		return []entity.IMoveGenerator{moves.NewKnightMoves()}
	case entity.Bishop:
		return []entity.IMoveGenerator{moves.NewDiagonal1Move(), moves.NewDiagonal2Move()}
	case entity.Rook:
		return []entity.IMoveGenerator{moves.NewHorizontalMove(), moves.NewVerticalMove()}
	case entity.Queen:
		return []entity.IMoveGenerator{moves.NewHorizontalMove(), moves.NewVerticalMove(), moves.NewDiagonal1Move(), moves.NewDiagonal2Move()}
	case entity.King:
		return []entity.IMoveGenerator{moves.KingMoves{}}
	}
	return nil
}

func GetPieceByType(pieceType entity.PieceType, color entity.Color) entity.Piece {
	switch pieceType {
	case entity.Pawn:
		return entity.Piece{Color: color, MoveGenerator: GetMoveGenerator(entity.Pawn)}
	case entity.Knight:
		return entity.Piece{Color: color, MoveGenerator: GetMoveGenerator(entity.Knight)}
	case entity.Bishop:
		return entity.Piece{Color: color, MoveGenerator: GetMoveGenerator(entity.Bishop)}
	case entity.Rook:
		return entity.Piece{Color: color, MoveGenerator: GetMoveGenerator(entity.Rook)}
	case entity.Queen:
		return entity.Piece{Color: color, MoveGenerator: GetMoveGenerator(entity.Queen)}
	case entity.King:
		return entity.Piece{Color: color, MoveGenerator: GetMoveGenerator(entity.King)}
	}
	return entity.Piece{}
}
