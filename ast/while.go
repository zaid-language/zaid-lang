package ast

import "zaidlang.org/x/zaid/token"

type While struct {
	ExpressionNode
	Token       token.Token
	Condition   ExpressionNode
	Consequence *Block
}
