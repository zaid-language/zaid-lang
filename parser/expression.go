package parser

import (
	"github.com/zaid-language/zaid-lang/ast"
	"github.com/zaid-language/zaid-lang/token"
)

func (parser *Parser) parseExpression(precedence int) ast.ExpressionNode {
	prefix := parser.prefixParserFns[parser.currentToken.Type]

	if prefix == nil {
		return nil
	}

	leftExpression := prefix()

	for !parser.nextTokenIs(token.SEMICOLON) && precedence < parser.nextTokenPrecedence() {
		infix := parser.infixParserFns[parser.nextToken.Type]

		if infix == nil {
			return leftExpression
		}

		parser.readToken()

		leftExpression = infix(leftExpression)
	}

	return leftExpression
}

func (parser *Parser) parseExpressionList(end token.Type) []ast.ExpressionNode {
	list := []ast.ExpressionNode{}

	if parser.nextTokenIs(end) {
		parser.readToken()

		return list
	}

	parser.readToken()

	list = append(list, parser.parseExpression(LOWEST))

	for parser.nextTokenIs(token.COMMA) {
		parser.readToken()
		parser.readToken()
		list = append(list, parser.parseExpression(LOWEST))
	}

	if !parser.expectNextTokenIs(end) {
		return nil
	}

	return list
}
