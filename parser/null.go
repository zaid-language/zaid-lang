package parser

import "zaidlang.org/x/zaid/ast"

func (parser *Parser) nullLiteral() ast.ExpressionNode {
	return &ast.Null{Token: parser.currentToken}
}
