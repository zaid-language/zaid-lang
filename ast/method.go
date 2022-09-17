package ast

import "github.com/zaid-language/zaid/token"

type Method struct {
	ExpressionNode
	Token     token.Token
	Left      ExpressionNode
	Method    ExpressionNode
	Arguments []ExpressionNode
}
