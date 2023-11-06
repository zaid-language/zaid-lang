package ast

import "zaidlang.org/x/zaid/token"

type Use struct {
	ExpressionNode
	Token  token.Token
	Traits []*Identifier
}