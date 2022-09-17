package ast

import "zaidlang.tech/x/zaid/token"

type Property struct {
	ExpressionNode
	AssignmentNode
	Token    token.Token
	Left     ExpressionNode
	Property ExpressionNode
}
