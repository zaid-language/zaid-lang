package ast

import "zaidlang.tech/x/zaid/token"

type Import struct {
	ExpressionNode
	Token token.Token
	Path  *String
}
