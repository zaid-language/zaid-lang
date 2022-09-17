package parser

import (
	"github.com/zaid-language/zaid/ast"
	"github.com/zaid-language/zaid/token"
)

func (parser *Parser) booleanLiteral() ast.ExpressionNode {
	return &ast.Boolean{
		Token: parser.currentToken,
		Value: parser.currentTokenIs(token.TRUE),
	}
}
