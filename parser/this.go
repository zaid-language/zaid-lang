package parser

import "zaidlang.org/x/zaid/ast"

func (parser *Parser) thisExpression() ast.ExpressionNode {
	return &ast.This{Token: parser.currentToken}
}
