package ast

import "github.com/zaid-language/zaid-lang/token"

type Break struct {
	ExpressionNode
	Token token.Token
}
