package ast

import "github.com/zaid-language/zaid-lang/token"

type Identifier struct {
	ExpressionNode
	AssignmentNode
	Token token.Token
	Value string
}
