package ast

import "zaidlang.tech/x/zaid/token"

type Assign struct {
	StatementNode
	Token token.Token
	Name  AssignmentNode
	Value ExpressionNode
}
