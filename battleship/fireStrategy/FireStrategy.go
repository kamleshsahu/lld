package fireStrategy

import "lld/battleship/entity"

type FireStrategy interface {
	Init(player []entity.Field)
	GetFireLocation(region int) (*entity.Cell, error)
}
