package parser

import (
	"zaidlang.org/x/zaid/ast"
	"zaidlang.org/x/zaid/token"
)

func (parser *Parser) useExpression() ast.ExpressionNode {
	use := &ast.Use{Token: parser.currentToken}

	if !parser.expectNextTokenIs(token.IDENTIFIER) {
		return nil
	}

	identifier := &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Lexeme}

	use.Traits = append(use.Traits, identifier)

	return use
}