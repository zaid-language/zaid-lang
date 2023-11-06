package ast

import "zaidlang.org/x/zaid/token"

type Trait struct {
	ExpressionNode
	Token token.Token
	Name  *Identifier
	Body  *Block
}