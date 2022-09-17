package ast

import "github.com/zaid-language/zaid-lang/token"

type Continue struct {
	ExpressionNode
	Token token.Token
}
