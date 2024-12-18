package eliminationStrategy

import "lld/battleship/entity"

type DefaultEliminationStrategy struct {
}

func (d *DefaultEliminationStrategy) IsEliminated(player *entity.Player) bool {
	return player.AllShipsKilled()
}

func NewDefaultEliminationStrategy() IEliminationStrategy {
	return &DefaultEliminationStrategy{}
}
