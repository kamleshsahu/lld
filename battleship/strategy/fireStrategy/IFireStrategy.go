package fireStrategy

import "lld/battleship/entity"

type FireStrategy interface {
	Init(playerFields []*entity.Field)
	GetFireLocation(playerFieldId int) (*entity.Cell, error)
}
