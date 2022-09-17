package ast

import "github.com/zaid-language/zaid/token"

type For struct {
	ExpressionNode
	Token       token.Token
	Identifier  *Identifier
	Initializer StatementNode
	Condition   ExpressionNode
	Increment   StatementNode
	Block       *Block
}
