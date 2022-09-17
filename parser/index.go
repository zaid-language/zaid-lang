package parser

import (
	"github.com/zaid-language/zaid/ast"
	"github.com/zaid-language/zaid/token"
)

func (parser *Parser) indexExpression(left ast.ExpressionNode) ast.ExpressionNode {
	expression := &ast.Index{Token: parser.currentToken, Left: left}

	parser.readToken()

	expression.Index = parser.parseExpression(LOWEST)

	if !parser.expectNextTokenIs(token.RIGHTBRACKET) {
		return nil
	}

	parser.previousIndex = expression
	parser.previousProperty = nil

	return expression
}
