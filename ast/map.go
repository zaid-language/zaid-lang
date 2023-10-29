package ast

import "zaidlang.org/x/zaid/token"

type Map struct {
	ExpressionNode
	Token token.Token
	Pairs map[ExpressionNode]ExpressionNode
}
