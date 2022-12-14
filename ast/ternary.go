package ast

import "github.com/zaid-language/zaid-lang/token"

type Ternary struct {
	ExpressionNode
	Token     token.Token
	Condition ExpressionNode
	IfTrue    ExpressionNode
	IfFalse   ExpressionNode
}
