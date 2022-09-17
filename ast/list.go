package ast

import "github.com/zaid-language/zaid/token"

type List struct {
	ExpressionNode
	Token    token.Token
	Elements []ExpressionNode
}
