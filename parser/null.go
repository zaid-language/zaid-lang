package parser

import "github.com/zaid-language/zaid-lang/ast"

func (parser *Parser) nullLiteral() ast.ExpressionNode {
	return &ast.Null{Token: parser.currentToken}
}
