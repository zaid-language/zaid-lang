package ast

import (
	"github.com/zaid-language/zaid/token"
)

type Function struct {
	ExpressionNode
	Token      token.Token
	Name       *Identifier
	Parameters []*Identifier
	Defaults   map[string]ExpressionNode
	Body       *Block
}
