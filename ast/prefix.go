package ast

import "zaidlang.org/x/zaid/token"

type Prefix struct {
	ExpressionNode
	Token    token.Token
	Operator string
	Right    ExpressionNode
}
