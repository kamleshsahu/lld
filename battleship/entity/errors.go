package entity

import (
	"fmt"
)

var (
	INVALID_CELL_OWNER      = "ship cant be placed at %s, %s is owned by %s"
	INVALID_CELL_ADD_SHIP   = "ship cant be placed at %s, 1 ship %s conflict"
	NO_CELL_LEFT            = "playerId : %d, no cell left to bomb, (probably all cells bombed)"
	NO_SHIP_PRESENT_IN_CELL = "no ship present at cell: %s"
)

func ErrInvalidCellOwner(cell Cell, cell2 Cell, player Player) error {
	return fmt.Errorf(INVALID_CELL_OWNER, cell.ToString(), cell2.ToString(), player.ToString())
}

func ErrNoCellLeft(playerId int) error {
	return fmt.Errorf(NO_CELL_LEFT, playerId)
}

func ErrInvalidCellShip(cell Cell, cell2 Cell) error {
	return fmt.Errorf(INVALID_CELL_ADD_SHIP, cell.ToString(), cell2.ToString())
}

func ErrNoShipPresentInCell(cell Cell) error {
	return fmt.Errorf(NO_SHIP_PRESENT_IN_CELL, cell.ToString())
}
