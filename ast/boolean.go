package ast

import "zaidlang.tech/x/zaid/token"

type Boolean struct {
	ExpressionNode
	Token token.Token
	Value bool
}
