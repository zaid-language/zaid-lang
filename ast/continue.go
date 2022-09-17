package ast

import "github.com/zaid-language/zaid/token"

type Continue struct {
	ExpressionNode
	Token token.Token
}
