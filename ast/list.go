package ast

import "github.com/zaid-language/zaid-lang/token"

type List struct {
	ExpressionNode
	Token    token.Token
	Elements []ExpressionNode
}
