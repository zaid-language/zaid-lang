package ast

import "zaidlang.tech/x/zaid/token"

type Map struct {
	ExpressionNode
	Token token.Token
	Pairs map[ExpressionNode]ExpressionNode
}
