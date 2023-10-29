package ast

import "zaidlang.org/x/zaid/token"

type Postfix struct {
	ExpressionNode
	Token    token.Token
	Operator string
}
