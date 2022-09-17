package parser

import "github.com/zaid-language/zaid/ast"

func (parser *Parser) stringLiteral() ast.ExpressionNode {
	return &ast.String{Token: parser.currentToken, Value: parser.currentToken.Literal.(string)}
}
