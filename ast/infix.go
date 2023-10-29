package ast

import "zaidlang.org/x/zaid/token"

type Infix struct {
	ExpressionNode
	Token    token.Token
	Left     ExpressionNode
	Operator string
	Right    ExpressionNode
}
