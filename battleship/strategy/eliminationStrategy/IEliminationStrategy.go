package eliminationStrategy

import "lld/battleship/entity"

type IEliminationStrategy interface {
	IsEliminated(*entity.Player) bool
}
