package utils

import (
	"fmt"
	"lld/battleship/entity"
	"strings"
)

func ClonePlayerFields(players []*entity.Player) []*entity.Field {
	playerField := make([]*entity.Field, len(players))
	for playerId, player := range players {
		cp := player.Field.Copy()
		playerField[playerId] = &cp
	}
	return playerField
}

func TurnMessage(currentPlayer, targetPlayer string, hitPosition string, ship *entity.Ship) string {
	msg := strings.Builder{}
	msg.WriteString(fmt.Sprintf("Player%s's turn: Missile fired at %s.", currentPlayer, hitPosition))
	if ship != nil {
		msg.WriteString(fmt.Sprintf("\"Hit\". Player%s's ship with id %s destroyed\n", targetPlayer, ship.GetId()))
	} else {
		msg.WriteString("\"Miss\"")
	}
	return msg.String()
}
