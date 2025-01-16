package parser

import "lld/cronParser/entity"

type IParser interface {
	Parse(token string, schedule *entity.Expression) string
}
