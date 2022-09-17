package ast

import "github.com/zaid-language/zaid-lang/token"

type Map struct {
	ExpressionNode
	Token token.Token
	Pairs map[ExpressionNode]ExpressionNode
}
