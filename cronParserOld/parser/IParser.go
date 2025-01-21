package parser

import "lld/cronParserOld/entity"

type IParser interface {
	Parse(token string, schedule *entity.Expression) string
}
