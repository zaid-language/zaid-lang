package ast

import "zaidlang.tech/x/zaid/token"

type Return struct {
	StatementNode
	Token token.Token
	Value ExpressionNode
}
