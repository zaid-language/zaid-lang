package parser

import (
	"zaidlang.org/x/zaid/ast"
	"zaidlang.org/x/zaid/token"
)

func (parser *Parser) whileExpression() ast.ExpressionNode {
	expression := &ast.While{Token: parser.currentToken}

	if !parser.expectNextTokenIs(token.LEFTPAREN) {
		return nil
	}

	parser.readToken()
	expression.Condition = parser.parseExpression(LOWEST)

	if !parser.expectNextTokenIs(token.RIGHTPAREN) {
		return nil
	}

	if !parser.expectNextTokenIs(token.LEFTBRACE) {
		return nil
	}

	expression.Consequence = parser.blockStatement()

	return expression
}
