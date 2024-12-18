package game

import "lld/battleship/entity"

type IGame interface {
	InitGame(N int) error
	AddShip(id string, size, player1ShipX, player1ShipY, player2ShipX, player2ShipY int) error
	StartGame() error
	ViewBattleField() string

	// just used for testing
	GetBoard() *entity.Board
	GetPlayers() []*entity.Player
	Fire(cell *entity.Cell) (*entity.Ship, error)
	Reset()
}
