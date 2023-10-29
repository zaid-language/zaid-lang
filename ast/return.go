package ast

import "zaidlang.org/x/zaid/token"

type Return struct {
	StatementNode
	Token token.Token
	Value ExpressionNode
}
