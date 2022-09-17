package ast

import "github.com/zaid-language/zaid-lang/token"

type Compound struct {
	ExpressionNode
	Token    token.Token
	Left     ExpressionNode
	Operator string
	Right    ExpressionNode
}
