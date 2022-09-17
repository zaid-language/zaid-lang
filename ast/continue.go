package ast

import "zaidlang.tech/x/zaid/token"

type Continue struct {
	ExpressionNode
	Token token.Token
}
