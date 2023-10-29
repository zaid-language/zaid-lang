package ast

import "zaidlang.org/x/zaid/token"

type Block struct {
	StatementNode
	Token      token.Token
	Statements []StatementNode
}
