package ast

import "github.com/zaid-language/zaid/token"

type Return struct {
	StatementNode
	Token token.Token
	Value ExpressionNode
}
