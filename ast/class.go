package ast

import "zaidlang.org/x/zaid/token"

type Class struct {
	ExpressionNode
	Token token.Token
	Name  *Identifier
	Super *Identifier
	Body  *Block
}
