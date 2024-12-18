package entity

type Ship struct {
	Id       string
	Location *Cell
	Size     int
	Owner    *Player
}

func NewShip(id string, location *Cell, size int, owner *Player) *Ship {
	return &Ship{id, location, size, owner}
}

func (ship *Ship) GetId() string {
	return ship.Id
}
