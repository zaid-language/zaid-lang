package parser

import "github.com/zaid-language/zaid-lang/ast"

func (parser *Parser) stringLiteral() ast.ExpressionNode {
	return &ast.String{Token: parser.currentToken, Value: parser.currentToken.Literal.(string)}
}
