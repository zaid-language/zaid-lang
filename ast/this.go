package ast

import "github.com/zaid-language/zaid-lang/token"

type This struct {
	ExpressionNode
	Token token.Token
}
