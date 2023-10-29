package ast

import "zaidlang.org/x/zaid/token"

type Boolean struct {
	ExpressionNode
	Token token.Token
	Value bool
}
