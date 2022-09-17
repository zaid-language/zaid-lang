package parser

import "github.com/zaid-language/zaid-lang/ast"

func (parser *Parser) breakStatement() ast.ExpressionNode {
	return &ast.Break{Token: parser.currentToken}
}
