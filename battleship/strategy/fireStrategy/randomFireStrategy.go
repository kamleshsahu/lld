package fireStrategy

import (
	"lld/battleship/entity"
	"math/rand"
	"sync"
)

type RandomFireStrategy struct {
	playerFields []*entity.Field
}

func (r *RandomFireStrategy) Init(playerField []*entity.Field) {
	r.playerFields = playerField
}

func (r *RandomFireStrategy) GetFireLocation(targetPlayerId int) (*entity.Cell, error) {
	field := r.playerFields[targetPlayerId]
	size := len(field.Cells)
	if size == 0 {
		return nil, entity.ErrNoCellLeft(targetPlayerId)
	}
	id := rand.Intn(size)
	ans := field.Cells[id]
	// remove that id from list
	field.Cells[id] = field.Cells[size-1]
	field.Cells = field.Cells[:size-1]
	return &ans, nil
}

var singleton sync.Once
var randomFireStrategy *RandomFireStrategy

func NewRandomFireStrategy() FireStrategy {
	singleton.Do(func() {
		randomFireStrategy = &RandomFireStrategy{}
	})
	return randomFireStrategy
}
