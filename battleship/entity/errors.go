package entity

import (
	"errors"
	"fmt"
)

var (
	INVALID_CELL_OWNER               = "ship cant be placed at %s, %s is owned by %s"
	INVALID_CELL_ADD_SHIP            = "ship cant be placed at %s, conflict with %s"
	NO_CELL_LEFT                     = "playerId : %d, no cell left to bomb, (probably all cells bombed)"
	NO_SHIP_PRESENT_IN_CELL          = "no ship present at cell: %s"
	BOARD_CAN_NOT_BE_DIVIDED_EQUALLY = "board cant be divided equally in %d / %d parts"
)

var (
	ERR_CELL_OUT_OF_BOUNDARY        = errors.New("cell out of boundary")
	ERR_GAME_HAS_NO_BOARD           = errors.New("game is not initialised")
	ERR_GAME_HAS_NO_PLAYERS         = errors.New("game is not initialised")
	ERR_GAME_HAS_LESS_THAN_2_PLAYER = errors.New("game has less than 2 players")
	ERR_GAME_ALREADY_STARTED        = errors.New("game already started")
)

func ErrInvalidCellOwner(cell *Cell, cell2 *Cell, player *Player) error {
	return fmt.Errorf(INVALID_CELL_OWNER, cell.ToString(), cell2.ToString(), player.ToString())
}

func ErrNoCellLeft(playerId int) error {
	return fmt.Errorf(NO_CELL_LEFT, playerId)
}

func ErrInvalidCellShip(cell *Cell, shipId string) error {
	return fmt.Errorf(INVALID_CELL_ADD_SHIP, cell.ToString(), shipId)
}

func ErrNoShipPresentInCell(cell *Cell) error {
	return fmt.Errorf(NO_SHIP_PRESENT_IN_CELL, cell.ToString())
}

func ErrEqualDivide(n int, playerCount int) error {
	return fmt.Errorf(BOARD_CAN_NOT_BE_DIVIDED_EQUALLY, n, playerCount)
}
