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
	for j := range cells {
		cells[j] = make([]*Cell, N)
		for i := range cells[j] {
			cells[j][i] = &Cell{X: i, Y: j}
		}
	}
	return &Board{cells: cells}
}

func (b *Board) AddShip(ship *Ship) error {
	size := ship.Size
	startLocation := ship.Location
	for i := startLocation.X; i < startLocation.X+size; i++ {
		for j := startLocation.Y; j < startLocation.Y+size; j++ {
			b.cells[j][i].Ship = ship
		}
	}
	fmt.Printf("ship added at startLocation %s\n", startLocation.ToString())
	return nil
}

func (b *Board) CanPlaceShip(startLocation *Cell, size int, owner *Player) (bool, error) {
	endLocation := NewCell(startLocation.X+size-1, startLocation.Y+size-1)
	if !b.IsValidCell(startLocation) || !b.IsValidCell(endLocation) {
		return false, ERR_CELL_OUT_OF_BOUNDARY
	}
	for i := startLocation.X; i < startLocation.X+size; i++ {
		for j := startLocation.Y; j < startLocation.Y+size; j++ {
			if b.cells[j][i].Ship != nil {
				return false, ErrInvalidCellShip(startLocation, b.cells[j][i].Ship.GetId())
			}
			if b.cells[j][i].Owner != owner {
				return false, ErrInvalidCellOwner(startLocation, b.cells[j][i], b.cells[j][i].Owner)
			}
		}
	}
	return true, nil
}

func (b *Board) RemoveShip(cell *Cell) (*Ship, error) {
	cell = b.GetCells()[cell.Y][cell.X]
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
	bf.WriteString("\n")
	bf.WriteString(fmt.Sprintf("%8s", " "))
	for colId, _ := range b.cells[0] {
		bf.WriteString(fmt.Sprintf("%8d", colId))
	}
	bf.WriteString("\n")
	for j, row := range b.GetCells() {
		bf.WriteString(fmt.Sprintf("%8d", j))
		for i, col := range row {
			cellOwner := ""
			if b.cells[j][i].Owner != nil {
				cellOwner = col.Owner.GetName()
			}
			if b.cells[j][i].Ship != nil {
				bf.WriteString(fmt.Sprintf("%8s", cellOwner+"-"+col.Ship.GetId()))
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
	return b.cells[cell.Y][cell.X].HasShip()
}

func (b *Board) GetCells() [][]*Cell {
	return b.cells
}
