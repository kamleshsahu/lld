package game

import (
	"errors"
	"lld/chess/entity"
)

type IGame interface {
	MovePiece(x, y int, x1, y1 int) (bool, error)
	isValidMove(fromCell *entity.Cell, toCell *entity.Cell) bool
	GetBoard() *entity.Board
}

type Game struct {
	Board entity.Board
}

func (g *Game) GetBoard() *entity.Board {
	return &g.Board
}

func NewGame() IGame {
	return &Game{entity.NewBoard()}
}

func (g *Game) MovePiece(x, y int, x1, y1 int) (bool, error) {

	fromCell, exist := g.Board.GetCell(x, y)
	if !exist || fromCell.Piece == nil {
		return false, errors.New("invalid move")
	}

	toCell, exist := g.Board.GetCell(x1, y1)
	if !exist || toCell.HasSamePieceType(fromCell.Piece) {
		return false, errors.New("invalid move")
	}

	if g.isValidMove(fromCell, toCell) {
		toCell.Piece = fromCell.Piece
		fromCell.Piece = nil
		return true, nil
	}

	return false, errors.New("invalid move")
}

func (g *Game) isValidMove(fromCell *entity.Cell, toCell *entity.Cell) bool {
	for _, generator := range fromCell.Piece.MoveGenerator {
		possibleCells := generator.GetPossibleMoves(g.Board, *fromCell)
		for _, cell := range possibleCells {
			if cell.Equals(toCell) {
				return true
			}
		}
	}
	return false
}
