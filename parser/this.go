package parser

import "zaidlang.tech/x/zaid/ast"

func (parser *Parser) thisExpression() ast.ExpressionNode {
	return &ast.This{Token: parser.currentToken}
}
