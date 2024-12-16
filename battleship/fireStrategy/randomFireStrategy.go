package fireStrategy

import (
	"lld/battleship/entity"
	"math/rand"
)

type RandomFireStrategy struct {
	players []entity.Field
}

func (r *RandomFireStrategy) Init(playerField []entity.Field) {
	r.players = playerField
}

func (r *RandomFireStrategy) GetFireLocation(playerID int) (*entity.Cell, error) {
	field := r.players[playerID]
	size := len(field.Cells)
	if size == 0 {
		return nil, entity.ErrNoCellLeft(playerID)
	}
	id := rand.Intn(size)
	ans := field.Cells[id]
	// remove that id from list
	field.Cells[id] = field.Cells[size-1]
	field.Cells = field.Cells[:size-1]
	return &ans, nil
}

func NewRandomFireStrategy() FireStrategy {
	return &RandomFireStrategy{}
}
