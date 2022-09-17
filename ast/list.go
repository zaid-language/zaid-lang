package ast

import "zaidlang.tech/x/zaid/token"

type List struct {
	ExpressionNode
	Token    token.Token
	Elements []ExpressionNode
}
