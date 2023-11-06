package parser

import (
	"zaidlang.org/x/zaid/ast"
	"zaidlang.org/x/zaid/token"
)

func (parser *Parser) traitStatement() ast.ExpressionNode {
	trait := &ast.Trait{Token: parser.currentToken}

	parser.readToken()

	trait.Name = &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Lexeme}

	if !parser.expectNextTokenIs(token.LEFTBRACE) {
		return nil
	}

	trait.Body = parser.blockStatement()

	return trait
}