package ast

import "github.com/zaid-language/zaid-lang/token"

type If struct {
	ExpressionNode
	Token       token.Token
	Condition   ExpressionNode
	Consequence *Block
	Alternative *Block
}
