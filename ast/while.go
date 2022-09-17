package ast

import "github.com/zaid-language/zaid/token"

type While struct {
	ExpressionNode
	Token       token.Token
	Condition   ExpressionNode
	Consequence *Block
}
