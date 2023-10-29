package ast

import "zaidlang.org/x/zaid/token"

type Index struct {
	ExpressionNode
	AssignmentNode
	Token token.Token
	Left  ExpressionNode
	Index ExpressionNode
}
