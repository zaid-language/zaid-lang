package parser

import (
	"github.com/zaid-language/zaid/ast"
	"github.com/zaid-language/zaid/token"
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
