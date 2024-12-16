package entity

type Player struct {
	Name      string
	Id        int
	ShipCount int
	Field     Field
}

func (p *Player) HasShip() bool {
	return p.ShipCount > 0
}

func (p *Player) ToString() string {
	return p.Name
}
