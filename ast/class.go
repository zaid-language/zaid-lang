package ast

import "zaidlang.tech/x/zaid/token"

type Class struct {
	ExpressionNode
	Token token.Token
	Name  *Identifier
	Super *Identifier
	Body  *Block
}
