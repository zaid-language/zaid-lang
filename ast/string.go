package ast

import (
	"zaidlang.org/x/zaid/token"
)

type String struct {
	ExpressionNode
	Token token.Token
	Value string
}
