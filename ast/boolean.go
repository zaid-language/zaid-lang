package ast

import "github.com/zaid-language/zaid/token"

type Boolean struct {
	ExpressionNode
	Token token.Token
	Value bool
}
