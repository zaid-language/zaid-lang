package ast

import "github.com/zaid-language/zaid-lang/token"

type Postfix struct {
	ExpressionNode
	Token    token.Token
	Operator string
}