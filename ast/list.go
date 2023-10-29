package ast

import "zaidlang.org/x/zaid/token"

type List struct {
	ExpressionNode
	Token    token.Token
	Elements []ExpressionNode
}
