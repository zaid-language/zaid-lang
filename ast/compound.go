package ast

import "zaidlang.org/x/zaid/token"

type Compound struct {
	ExpressionNode
	Token    token.Token
	Left     ExpressionNode
	Operator string
	Right    ExpressionNode
}
