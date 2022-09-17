package ast

import "github.com/zaid-language/zaid-lang/token"

type Assign struct {
	StatementNode
	Token token.Token
	Name  AssignmentNode
	Value ExpressionNode
}
