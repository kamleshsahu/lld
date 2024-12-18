package entity

type Player struct {
	Name      string
	Id        int
	ShipCount int
	Field     Field
}

func (p *Player) AllShipsKilled() bool {
	return p.ShipCount > 0
}

func (p *Player) GetName() string {
	if p.Name == "" {
		return string(byte(p.Id + 'A'))
	}
	return p.Name
}

func (p *Player) ToString() string {
	return p.Name
}

func NewPlayer(id int) *Player {
	return &Player{Id: id}
}
