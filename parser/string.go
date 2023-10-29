package parser

import "zaidlang.org/x/zaid/ast"

func (parser *Parser) stringLiteral() ast.ExpressionNode {
	return &ast.String{Token: parser.currentToken, Value: parser.currentToken.Literal.(string)}
}
