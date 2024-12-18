package board

import (
	"errors"
	"lld/battleship/entity"
	"testing"
)

func TestNewBoard(t *testing.T) {
	N := 5
	board := entity.NewBoard(N)

	if len(board.GetCells()) != N {
		t.Errorf("Expected board size %d, got %d", N, len(board.GetCells()))
	}

	for i := range board.GetCells() {
		if len(board.GetCells()[i]) != N {
			t.Errorf("Expected row size %d, got %d", N, len(board.GetCells()[i]))
		}
	}
}

func TestAddShip(t *testing.T) {
	N := 5
	board := entity.NewBoard(N)
	ship := &entity.Ship{Location: &entity.Cell{X: 1, Y: 1}, Size: 2}

	err := board.AddShip(ship)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	startLocation := ship.Location
	for i := startLocation.X; i < startLocation.X+ship.Size; i++ {
		for j := startLocation.Y; j < startLocation.Y+ship.Size; j++ {
			if board.GetCells()[j][i].Ship != ship {
				t.Errorf("Expected ship at cell (%d, %d), but found none", j, i)
			}
		}
	}
}

func TestCanPlaceShip(t *testing.T) {
	N := 5
	board := entity.NewBoard(N)
	startLocation1 := &entity.Cell{X: 1, Y: 1}
	size := 2

	// Valid location
	valid, err := board.CanPlaceShip(startLocation1, size, nil)
	if !valid || err != nil {
		t.Errorf("Expected location to be valid, but got valid=%v, err=%v", valid, err)
	}

	// Out of bound location
	startLocation2 := &entity.Cell{X: 4, Y: 4}
	valid, err = board.CanPlaceShip(startLocation2, size, nil)
	if valid {
		t.Errorf("Expected location to be invalid, but got valid=true")
	}

	if !errors.Is(err, entity.ERR_CELL_OUT_OF_BOUNDARY) {
		t.Errorf("Expected error %v, but got %v", entity.ERR_CELL_OUT_OF_BOUNDARY, err)
	}

	p1 := entity.NewPlayer(1)
	err = board.AddShip(entity.NewShip("SH1", startLocation1, size, p1))
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Ship conflict
	startLocation2 = &entity.Cell{X: 2, Y: 2}
	valid, err = board.CanPlaceShip(startLocation2, size, nil)
	if valid {
		t.Errorf("Expected location to be invalid, but got valid=true")
	}

	expected := entity.ErrInvalidCellShip(startLocation2, "SH1")
	if err.Error() != expected.Error() {
		t.Errorf("Expected error message %q, but got %q", expected.Error(), err.Error())
	}

}

func TestRemoveShip(t *testing.T) {
	// Create a new board
	boardSize := 5
	board := entity.NewBoard(boardSize)

	// Define a ship
	ship := &entity.Ship{
		Location: &entity.Cell{X: 1, Y: 1},
		Size:     2, // Ship occupies a 2x2 square
	}

	// Add the ship to the board
	err := board.AddShip(ship)
	if err != nil {
		t.Fatalf("Failed to add ship: %v", err)
	}

	// Attempt to remove the ship
	removedShip, err := board.RemoveShip(&entity.Cell{X: 1, Y: 1})
	if err != nil {
		t.Errorf("Failed to remove ship: %v", err)
	}

	// Verify the returned ship is correct
	if removedShip != ship {
		t.Errorf("Expected removed ship %v, got %v", ship, removedShip)
	}

	// Verify the cells are cleared
	for i := 1; i < 1+ship.Size; i++ {
		for j := 1; j < 1+ship.Size; j++ {
			if board.GetCells()[j][i].Ship != nil {
				t.Errorf("Expected cell (%d, %d) to be cleared, but still contains a ship", i, j)
			}
		}
	}

	// Test removing a ship from an empty cell
	_, err = board.RemoveShip(&entity.Cell{X: 0, Y: 0})
	if err == nil {
		t.Errorf("Expected error when removing ship from an empty cell, but got nil")
	}

	// Verify the error message
	expected := entity.ErrNoShipPresentInCell(board.GetCells()[0][0])
	if err.Error() != expected.Error() {
		t.Errorf("Expected error %v, but got %v", expected.Error(), err.Error())
	}
}

func TestViewBattleField(t *testing.T) {
	// Create a new board of size 5x5
	boardSize := 5
	board := entity.NewBoard(boardSize)

	// Create ship objects with some starting locations
	ship1 := &entity.Ship{
		Id:       "SH1",
		Location: &entity.Cell{X: 1, Y: 1},
		Size:     2, // Ship occupies a 2x2 space
	}
	ship2 := &entity.Ship{
		Id:       "SH1",
		Location: &entity.Cell{X: 3, Y: 4},
		Size:     1, // Ship occupies a single cell
	}

	// Add the ships to the board
	if err := board.AddShip(ship1); err != nil {
		t.Fatalf("Failed to add ship1: %v", err)
	}
	if err := board.AddShip(ship2); err != nil {
		t.Fatalf("Failed to add ship2: %v", err)
	}

	t.Log(board.ViewBattleField())

}
