package entity

import (
	"errors"
	"fmt"
	"strings"
)

type Board struct {
	Cells [][]Cell
}

func NewBoard(N int) *Board {
	cells := make([][]Cell, N)
	for i := range cells {
		cells[i] = make([]Cell, N)
		for j := range cells[i] {
			cells[i][j] = Cell{X: i, Y: j}
		}
	}
	return &Board{Cells: cells}
}

func (b *Board) AddShip(ship *Ship, location Cell, size int) error {
	if isEmpty, err := b.isEmpty(location, size); !isEmpty {
		return err
	}
	for i := location.X; i < location.X+size; i++ {
		for j := location.Y; j < location.Y+size; j++ {
			b.Cells[i][j].Ship = ship
		}
	}
	fmt.Printf("ship added at location %s\n", location.ToString())
	return nil
}

func (b *Board) isEmpty(location Cell, size int) (bool, error) {
	owner := b.Cells[location.X][location.Y].Owner
	for i := location.X; i < location.X+size; i++ {
		for j := location.Y; j < location.Y+size; j++ {
			if b.Cells[i][j].Ship != nil {
				return false, ErrInvalidCellShip(location, b.Cells[i][j])
			}
			if b.Cells[i][j].Owner != owner {
				return false, ErrInvalidCellOwner(location, b.Cells[i][j], *b.Cells[i][j].Owner)
			}
		}
	}
	return true, nil
}

func (b *Board) RemoveShip(cell Cell) (*Ship, error) {
	if b.Cells[cell.X][cell.Y].Ship == nil {
		return nil, errors.New(fmt.Sprintf("no ship present at cell:", cell.ToString()))
	}
	location := cell.Ship.Location
	size := cell.Ship.Size
	for i := location.X; i < location.X+size; i++ {
		for j := location.Y; j < location.Y+size; j++ {
			b.Cells[i][j].Ship = nil
		}
	}
	return cell.Ship, nil
}

func (b *Board) ViewBattleField() string {
	bf := strings.Builder{}

	for i := 0; i < len(b.Cells); i++ {
		for j := 0; j < len(b.Cells[0]); j++ {
			if b.Cells[i][j].Ship != nil {
				bf.WriteString(fmt.Sprintf("%8s", string(b.Cells[i][j].Owner.Name[0])+"-"+b.Cells[i][j].Ship.Name))
			} else {
				bf.WriteString(fmt.Sprintf("%8s", "."))
			}
		}
		bf.WriteString("\n\n")
	}
	return bf.String()
}

func (b *Board) HasShip(location Cell) bool {
	return b.Cells[location.X][location.Y].Ship != nil
}
