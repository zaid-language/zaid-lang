package ast

import "github.com/zaid-language/zaid-lang/token"

type Return struct {
	StatementNode
	Token token.Token
	Value ExpressionNode
}
