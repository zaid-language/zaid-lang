package ast

import "zaidlang.org/x/zaid/token"

type Method struct {
	ExpressionNode
	Token     token.Token
	Left      ExpressionNode
	Method    ExpressionNode
	Arguments []ExpressionNode
}
