package ast

import "zaidlang.tech/x/zaid/token"

type Ternary struct {
	ExpressionNode
	Token     token.Token
	Condition ExpressionNode
	IfTrue    ExpressionNode
	IfFalse   ExpressionNode
}
