package ast

import "zaidlang.tech/x/zaid/token"

type While struct {
	ExpressionNode
	Token       token.Token
	Condition   ExpressionNode
	Consequence *Block
}
