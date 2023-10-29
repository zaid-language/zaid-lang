package ast

import "zaidlang.org/x/zaid/token"

type Continue struct {
	ExpressionNode
	Token token.Token
}
