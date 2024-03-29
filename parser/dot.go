package parser

import (
	"zaidlang.org/x/zaid/ast"
	"zaidlang.org/x/zaid/token"
)

func (parser *Parser) dotExpression(left ast.ExpressionNode) ast.ExpressionNode {
	currentToken := parser.currentToken
	currentPrecedence := parser.currentTokenPrecedence()

	parser.readToken()

	if parser.nextTokenIs(token.LEFTPAREN) {
		// Method
		expression := &ast.Method{Token: currentToken, Left: left}
		expression.Method = parser.parseExpression(currentPrecedence)

		parser.readToken()

		expression.Arguments = parser.parseExpressionList(token.RIGHTPAREN)

		return expression
	}

	// Property
	expression := &ast.Property{Token: currentToken, Left: left}
	expression.Property = parser.parseExpression(currentPrecedence)

	parser.previousProperty = expression
	parser.previousIndex = nil

	return expression
}
