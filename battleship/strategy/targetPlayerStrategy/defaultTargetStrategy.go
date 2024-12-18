package targetPlayerStrategy

import (
	"errors"
	"lld/battleship/entity"
)

type defaultTargetStrategy struct {
}

func (d defaultTargetStrategy) GetTargetPlayer(currentPlayer *entity.Player, allPlayers []*entity.Player) (*entity.Player, error) {
	if len(allPlayers) <= 0 {
		return nil, errors.New("game has less than 2 player")
	}

	targetPlayerId := (currentPlayer.Id + 1) % len(allPlayers)
	return allPlayers[targetPlayerId], nil
}

func NewDefaultTargetStrategy() ITargetPlayerStrategy {
	return &defaultTargetStrategy{}
}
