package ast

import "zaidlang.tech/x/zaid/token"

type Index struct {
	ExpressionNode
	AssignmentNode
	Token token.Token
	Left  ExpressionNode
	Index ExpressionNode
}
