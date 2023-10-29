package ast

import "zaidlang.org/x/zaid/token"

type Assign struct {
	StatementNode
	Token token.Token
	Name  AssignmentNode
	Value ExpressionNode
}
