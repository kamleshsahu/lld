package entity

import "fmt"

type Cell struct {
	X     int
	Y     int
	id    int
	Ship  *Ship
	Owner *Player
}

func (l *Cell) Copy() Cell {
	return Cell{X: l.X, Y: l.Y}
}

func (l *Cell) ToString() string {
	return fmt.Sprintf("%d#%d", l.X, l.Y)
}
