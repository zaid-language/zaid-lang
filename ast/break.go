package ast

import "zaidlang.tech/x/zaid/token"

type Break struct {
	ExpressionNode
	Token token.Token
}
