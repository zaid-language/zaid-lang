package parser

import "github.com/zaid-language/zaid/ast"

func (parser *Parser) nullLiteral() ast.ExpressionNode {
	return &ast.Null{Token: parser.currentToken}
}
