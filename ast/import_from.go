package ast

import "github.com/zaid-language/zaid-lang/token"

type ImportFrom struct {
	ExpressionNode
	Token       token.Token
	Path        *String
	Identifiers map[string]*Identifier
	Everything  bool
}
