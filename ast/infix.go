package ast

import "github.com/zaid-language/zaid/token"

type Infix struct {
	ExpressionNode
	Token    token.Token
	Left     ExpressionNode
	Operator string
	Right    ExpressionNode
}
