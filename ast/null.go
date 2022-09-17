package ast

import "github.com/zaid-language/zaid/token"

type Null struct {
	ExpressionNode
	Token token.Token
}
