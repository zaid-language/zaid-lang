package ast

import "zaidlang.tech/x/zaid/token"

type Null struct {
	ExpressionNode
	Token token.Token
}
