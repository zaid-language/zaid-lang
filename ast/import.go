package ast

import "github.com/zaid-language/zaid-lang/token"

type Import struct {
	ExpressionNode
	Token token.Token
	Path  *String
}
