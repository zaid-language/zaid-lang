package ast

import (
	"github.com/zaid-language/zaid-lang/token"
)

type String struct {
	ExpressionNode
	Token token.Token
	Value string
}
