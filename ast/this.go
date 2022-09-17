package ast

import "zaidlang.tech/x/zaid/token"

type This struct {
	ExpressionNode
	Token token.Token
}
