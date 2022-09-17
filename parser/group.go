package parser

import (
	"zaidlang.tech/x/zaid/ast"
	"zaidlang.tech/x/zaid/token"
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
