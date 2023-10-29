package ast

import "zaidlang.org/x/zaid/token"

type Break struct {
	ExpressionNode
	Token token.Token
}
