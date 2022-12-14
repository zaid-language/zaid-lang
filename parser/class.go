package parser

import (
	"github.com/zaid-language/zaid-lang/ast"
	"github.com/zaid-language/zaid-lang/token"
)

func (parser *Parser) classStatement() ast.ExpressionNode {
	class := &ast.Class{Token: parser.currentToken}

	parser.readToken()

	class.Name = &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Lexeme}

	if parser.nextTokenIs(token.EXTENDS) {
		parser.readToken()

		if !parser.expectNextTokenIs(token.IDENTIFIER) {
			return nil
		}

		class.Super = &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Lexeme}
	}

	if !parser.expectNextTokenIs(token.LEFTBRACE) {
		return nil
	}

	class.Body = parser.blockStatement()

	return class
}
