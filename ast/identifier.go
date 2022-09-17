package ast

import "zaidlang.tech/x/zaid/token"

type Identifier struct {
	ExpressionNode
	AssignmentNode
	Token token.Token
	Value string
}
