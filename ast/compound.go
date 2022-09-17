package ast

import "github.com/zaid-language/zaid/token"

type Compound struct {
	ExpressionNode
	Token    token.Token
	Left     ExpressionNode
	Operator string
	Right    ExpressionNode
}
