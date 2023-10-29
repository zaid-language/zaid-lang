package parser

import (
	"zaidlang.org/x/zaid/ast"
	"zaidlang.org/x/zaid/token"
)

func (parser *Parser) listLiteral() ast.ExpressionNode {
	list := &ast.List{Token: parser.currentToken}

	list.Elements = parser.parseExpressionList(token.RIGHTBRACKET)

	return list
}
