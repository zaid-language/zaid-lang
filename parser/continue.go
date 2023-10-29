package parser

import "zaidlang.org/x/zaid/ast"

func (parser *Parser) continueStatement() ast.ExpressionNode {
	return &ast.Continue{Token: parser.currentToken}
}
