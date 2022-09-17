package ast

import "github.com/zaid-language/zaid/token"

type Import struct {
	ExpressionNode
	Token token.Token
	Path  *String
}
