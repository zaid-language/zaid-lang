package ast

import "zaidlang.org/x/zaid/token"

type Ternary struct {
	ExpressionNode
	Token     token.Token
	Condition ExpressionNode
	IfTrue    ExpressionNode
	IfFalse   ExpressionNode
}
