package targetPlayerStrategy

import "lld/battleship/entity"

type ITargetPlayerStrategy interface {
	GetTargetPlayer(currentPlayer *entity.Player, allPlayers []*entity.Player) (*entity.Player, error)
}
