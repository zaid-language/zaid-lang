package ast

import "zaidlang.org/x/zaid/token"

type Import struct {
	ExpressionNode
	Token token.Token
	Path  *String
}
