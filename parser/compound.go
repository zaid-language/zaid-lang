package parser

import (
	"github.com/zaid-language/zaid/ast"
)

func (parser *Parser) compoundExpression(left ast.ExpressionNode) ast.ExpressionNode {
	compound := &ast.Compound{
		Token:    parser.currentToken,
		Operator: parser.currentToken.Lexeme,
		Left:     left,
	}

	precedence := parser.currentTokenPrecedence()

	parser.readToken()

	compound.Right = parser.parseExpression(precedence)

	return compound
}
