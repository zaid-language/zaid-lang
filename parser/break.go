package parser

import "github.com/zaid-language/zaid/ast"

func (parser *Parser) breakStatement() ast.ExpressionNode {
	return &ast.Break{Token: parser.currentToken}
}
