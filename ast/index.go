package ast

import "github.com/zaid-language/zaid/token"

type Index struct {
	ExpressionNode
	AssignmentNode
	Token token.Token
	Left  ExpressionNode
	Index ExpressionNode
}
