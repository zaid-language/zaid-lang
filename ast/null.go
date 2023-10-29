package ast

import "zaidlang.org/x/zaid/token"

type Null struct {
	ExpressionNode
	Token token.Token
}
