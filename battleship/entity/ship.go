package entity

type Ship struct {
	id       int
	Location Cell
	Name     string
	Size     int
	Owner    *Player
}
