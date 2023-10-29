package ast

import "zaidlang.org/x/zaid/token"

type ImportFrom struct {
	ExpressionNode
	Token       token.Token
	Path        *String
	Identifiers map[string]*Identifier
	Everything  bool
}
