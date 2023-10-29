package ast

import "zaidlang.org/x/zaid/token"

type This struct {
	ExpressionNode
	Token token.Token
}
