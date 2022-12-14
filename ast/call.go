package ast

import "github.com/zaid-language/zaid-lang/token"

type Call struct {
	ExpressionNode
	Token     token.Token
	Callee    ExpressionNode
	Arguments []ExpressionNode
}
