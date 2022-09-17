package ast

import "github.com/zaid-language/zaid-lang/token"

type Prefix struct {
	ExpressionNode
	Token    token.Token
	Operator string
	Right    ExpressionNode
}
