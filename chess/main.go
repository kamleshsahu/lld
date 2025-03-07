package main

import (
	"fmt"
	"lld/chess/entity"
	"lld/chess/factory"
	game2 "lld/chess/game"
)

func main() {

	//whitepawn := GetPieceByType(Pawn, White)

	game := game2.NewGame()

	//cell, _ := game.Board.GetCell(0, 0)
	//cell.Piece = &whitepawn
	//
	//moved, err := game.MovePiece(0, 0, 0, 2)
	//
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(moved)
	//}

	//whiteRook := GetPieceByType(Rook, White)
	//cell, _ := game.Board.GetCell(0, 0)
	//cell.Piece = &whiteRook
	//
	//moved, err := game.MovePiece(0, 0, 0, 6)
	//
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(moved)
	//}
	//
	//moved, err = game.MovePiece(0, 6, 0, 8)
	//
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(moved)
	//}
	board := game.GetBoard()

	whiteQueen := factory.GetPieceByType(entity.Queen, entity.White, *board)
	cell, _ := board.GetCell(4, 4)
	cell.Piece = &whiteQueen
	moved, err := game.MovePiece(4, 4, 1, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(moved)
	}

	pawn := factory.GetPieceByType(entity.Pawn, entity.White, *board)
	cell, _ = board.GetCell(1, 1)
	cell.Piece = &pawn
	moved, err = game.MovePiece(cell.X, cell.Y, 1, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(moved)
	}

	moved, err = game.MovePiece(1, 3, 1, 4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(moved)
	}

	bpawn := factory.GetPieceByType(entity.Pawn, entity.Black, *board)
	cell, _ = board.GetCell(1, 7)
	cell.Piece = &bpawn
	moved, err = game.MovePiece(cell.X, cell.Y, 1, 6)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(moved)
	}

	moved, err = game.MovePiece(1, 6, 1, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(moved)
	}
	board.PrintBoard()

	bknight := factory.GetPieceByType(entity.Knight, entity.Black, *board)
	cell, _ = board.GetCell(6, 7)
	cell.Piece = &bknight
	moved, err = game.MovePiece(cell.X, cell.Y, 5, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(moved)
	}

	board.PrintBoard()
}
