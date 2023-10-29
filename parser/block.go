package parser

import (
	"zaidlang.org/x/zaid/ast"
	"zaidlang.org/x/zaid/token"
)

func (parser *Parser) blockStatement() *ast.Block {
	block := &ast.Block{Token: parser.currentToken}
	block.Statements = []ast.StatementNode{}

	parser.readToken()

	for !parser.currentTokenIs(token.RIGHTBRACE) && !parser.isAtEnd() {
		statement := parser.statement()

		block.Statements = append(block.Statements, statement)

		parser.readToken()
	}

	return block
}
