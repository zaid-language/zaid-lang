package ast

import "zaidlang.org/x/zaid/token"

type Property struct {
	ExpressionNode
	AssignmentNode
	Token    token.Token
	Left     ExpressionNode
	Property ExpressionNode
}
