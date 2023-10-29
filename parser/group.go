package parser

import (
	"zaidlang.org/x/zaid/ast"
	"zaidlang.org/x/zaid/token"
)

func (parser *Parser) groupExpression() ast.ExpressionNode {
	// Read the opening token.LEFTPAREN ("(")
	parser.readToken()

	group := parser.parseExpression(LOWEST)

	if !parser.expectNextTokenIs(token.RIGHTPAREN) {
		return nil
	}

	return group
}
