package parser

import "zaidlang.org/x/zaid/ast"

func (parser *Parser) breakStatement() ast.ExpressionNode {
	return &ast.Break{Token: parser.currentToken}
}
