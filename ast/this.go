package ast

import "github.com/zaid-language/zaid/token"

type This struct {
	ExpressionNode
	Token token.Token
}
