package ast

import "github.com/zaid-language/zaid-lang/token"

type Method struct {
	ExpressionNode
	Token     token.Token
	Left      ExpressionNode
	Method    ExpressionNode
	Arguments []ExpressionNode
}
