package ast

import "github.com/zaid-language/zaid/token"

type Block struct {
	StatementNode
	Token      token.Token
	Statements []StatementNode
}
