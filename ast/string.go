package ast

import (
	"zaidlang.tech/x/zaid/token"
)

type String struct {
	ExpressionNode
	Token token.Token
	Value string
}
