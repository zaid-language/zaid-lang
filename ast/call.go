package ast

import "zaidlang.org/x/zaid/token"

type Call struct {
	ExpressionNode
	Token     token.Token
	Callee    ExpressionNode
	Arguments []ExpressionNode
}
