package ast

import "github.com/zaid-language/zaid/token"

type Property struct {
	ExpressionNode
	AssignmentNode
	Token    token.Token
	Left     ExpressionNode
	Property ExpressionNode
}
