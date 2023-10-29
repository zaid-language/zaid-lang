package ast

import "zaidlang.org/x/zaid/token"

type Identifier struct {
	ExpressionNode
	AssignmentNode
	Token token.Token
	Value string
}
