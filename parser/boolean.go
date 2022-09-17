package parser

import (
	"zaidlang.tech/x/zaid/ast"
	"zaidlang.tech/x/zaid/token"
)

func (parser *Parser) booleanLiteral() ast.ExpressionNode {
	return &ast.Boolean{
		Token: parser.currentToken,
		Value: parser.currentTokenIs(token.TRUE),
	}
}
