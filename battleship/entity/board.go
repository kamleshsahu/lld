package entity

import (
	"fmt"
	"strings"
)

type Board struct {
	cells [][]*Cell
}

func NewBoard(N int) *Board {
	cells := make([][]*Cell, N)
	for i := range cells {
		cells[i] = make([]*Cell, N)
		for j := range cells[i] {
			cells[i][j] = &Cell{X: i, Y: j}
		}
	}
	return &Board{cells: cells}
}

func (b *Board) AddShip(ship *Ship, startLocation *Cell, size int) error {
	for i := startLocation.X; i < startLocation.X+size; i++ {
		for j := startLocation.Y; j < startLocation.Y+size; j++ {
			b.cells[j][i].Ship = ship
		}
	}
	fmt.Printf("ship added at startLocation %s\n", startLocation.ToString())
	return nil
}

func (b *Board) IsValidLocation(startLocation *Cell, size int) (bool, error) {
	if !b.IsValidCell(startLocation) {
		return false, ERR_CELL_OUT_OF_BOUNDARY
	}
	owner := b.cells[startLocation.X][startLocation.Y].Owner
	for i := startLocation.X; i < startLocation.X+size; i++ {
		for j := startLocation.Y; j < startLocation.Y+size; j++ {
			if b.cells[i][j].Ship != nil {
				return false, ErrInvalidCellShip(startLocation, b.cells[i][j])
			}
			if b.cells[i][j].Owner != owner {
				return false, ErrInvalidCellOwner(startLocation, b.cells[i][j], b.cells[i][j].Owner)
			}
		}
	}
	return true, nil
}

func (b *Board) RemoveShip(cell *Cell) (*Ship, error) {
	cell = b.GetCells()[cell.X][cell.Y]
	if !cell.HasShip() {
		return nil, ErrNoShipPresentInCell(cell)
	}
	ship := cell.Ship
	startLocation := ship.Location
	size := ship.Size
	for i := startLocation.X; i < startLocation.X+size; i++ {
		for j := startLocation.Y; j < startLocation.Y+size; j++ {
			b.cells[j][i].Ship = nil
		}
	}
	return ship, nil
}

func (b *Board) ViewBattleField() string {
	bf := strings.Builder{}

	for i := 0; i < len(b.cells); i++ {
		for j := 0; j < len(b.cells[0]); j++ {
			if b.cells[i][j].Ship != nil {
				bf.WriteString(fmt.Sprintf("%8s", b.cells[i][j].Owner.GetName()+"-"+b.cells[i][j].Ship.GetId()))
			} else {
				bf.WriteString(fmt.Sprintf("%8s", "."))
			}
		}
		bf.WriteString("\n")
	}
	return bf.String()
}

func (b *Board) IsValidCell(cell *Cell) bool {
	if cell.X < 0 || cell.X >= len(b.cells[0]) || cell.Y < 0 || cell.Y >= len(b.cells) {
		return false
	}
	return true
}

func (b *Board) HasShip(cell *Cell) bool {
	return b.cells[cell.X][cell.Y].HasShip()
}

func (b *Board) GetCells() [][]*Cell {
	return b.cells
}
