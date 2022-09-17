package ast

import "github.com/zaid-language/zaid/token"

type Class struct {
	ExpressionNode
	Token token.Token
	Name  *Identifier
	Super *Identifier
	Body  *Block
}
