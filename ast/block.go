package ast

import "zaidlang.tech/x/zaid/token"

type Block struct {
	StatementNode
	Token      token.Token
	Statements []StatementNode
}
