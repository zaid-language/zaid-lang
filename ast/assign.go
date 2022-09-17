package ast

import "github.com/zaid-language/zaid/token"

type Assign struct {
	StatementNode
	Token token.Token
	Name  AssignmentNode
	Value ExpressionNode
}
