package ast

import "github.com/zaid-language/zaid-lang/token"

type Null struct {
	ExpressionNode
	Token token.Token
}
