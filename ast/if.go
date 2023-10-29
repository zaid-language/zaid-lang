package ast

import "zaidlang.org/x/zaid/token"

type If struct {
	ExpressionNode
	Token       token.Token
	Condition   ExpressionNode
	Consequence *Block
	Alternative *Block
}
