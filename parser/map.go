package parser

import (
	"github.com/zaid-language/zaid/ast"
	"github.com/zaid-language/zaid/token"
)

func (parser *Parser) mapLiteral() ast.ExpressionNode {
	mapLiteral := &ast.Map{Token: parser.currentToken}
	mapLiteral.Pairs = make(map[ast.ExpressionNode]ast.ExpressionNode)

	for !parser.nextTokenIs(token.RIGHTBRACE) {
		parser.readToken()

		key := parser.parseExpression(LOWEST)

		if !parser.expectNextTokenIs(token.COLON) {
			return nil
		}

		parser.readToken()

		value := parser.parseExpression(LOWEST)

		mapLiteral.Pairs[key] = value

		if !parser.nextTokenIs(token.RIGHTBRACE) && !parser.expectNextTokenIs(token.COMMA) {
			return nil
		}
	}

	if !parser.expectNextTokenIs(token.RIGHTBRACE) {
		return nil
	}

	return mapLiteral
}
