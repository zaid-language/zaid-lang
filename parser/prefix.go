package parser

import "github.com/zaid-language/zaid/ast"

func (parser *Parser) prefixExpression() ast.ExpressionNode {
	prefix := &ast.Prefix{
		Token:    parser.currentToken,
		Operator: parser.currentToken.Lexeme,
	}

	parser.readToken()

	prefix.Right = parser.parseExpression(PREFIX)

	return prefix
}
