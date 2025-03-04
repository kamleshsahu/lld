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

	whiteQueen := factory.GetPieceByType(entity.Queen, entity.White)
	board := game.GetBoard()
	cell, _ := board.GetCell(4, 4)
	cell.Piece = &whiteQueen
	moved, err := game.MovePiece(4, 4, 1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(moved)
	}

}
