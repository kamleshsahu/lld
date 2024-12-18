package targetPlayerStrategy

import (
	"lld/battleship/entity"
)

type defaultTargetStrategy struct {
}

func (d defaultTargetStrategy) GetTargetPlayer(currentPlayer *entity.Player, allPlayers []*entity.Player) (*entity.Player, error) {
	if len(allPlayers) <= 0 {
		return nil, entity.ERR_GAME_HAS_LESS_THAN_2_PLAYER
	}

	targetPlayerId := (currentPlayer.Id + 1) % len(allPlayers)
	return allPlayers[targetPlayerId], nil
}

func NewDefaultTargetStrategy() ITargetPlayerStrategy {
	return &defaultTargetStrategy{}
}
