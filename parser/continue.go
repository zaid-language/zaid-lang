package parser

import "github.com/zaid-language/zaid-lang/ast"

func (parser *Parser) continueStatement() ast.ExpressionNode {
	return &ast.Continue{Token: parser.currentToken}
}
