package ast

import "zaidlang.tech/x/zaid/token"

type ImportFrom struct {
	ExpressionNode
	Token       token.Token
	Path        *String
	Identifiers map[string]*Identifier
	Everything  bool
}
