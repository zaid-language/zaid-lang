package ast

import "zaidlang.tech/x/zaid/token"

type Prefix struct {
	ExpressionNode
	Token    token.Token
	Operator string
	Right    ExpressionNode
}
