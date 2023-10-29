package parser

import (
	"zaidlang.org/x/zaid/ast"
	"zaidlang.org/x/zaid/token"
)

func (parser *Parser) booleanLiteral() ast.ExpressionNode {
	return &ast.Boolean{
		Token: parser.currentToken,
		Value: parser.currentTokenIs(token.TRUE),
	}
}
