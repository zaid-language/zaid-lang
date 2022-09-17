package parser

import "github.com/zaid-language/zaid-lang/ast"

func (parser *Parser) thisExpression() ast.ExpressionNode {
	return &ast.This{Token: parser.currentToken}
}
