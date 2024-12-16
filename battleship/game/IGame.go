package game

import "lld/battleship/entity"

type IGame interface {
	InitGame(N int, players []entity.Player) error
	AddShip(ship *entity.Ship, location entity.Cell, Size int) error
	StartGame() error
}
