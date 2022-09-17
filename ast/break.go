package ast

import "github.com/zaid-language/zaid/token"

type Break struct {
	ExpressionNode
	Token token.Token
}
