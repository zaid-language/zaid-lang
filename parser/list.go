package parser

import (
	"zaidlang.tech/x/zaid/ast"
	"zaidlang.tech/x/zaid/token"
)

func (parser *Parser) listLiteral() ast.ExpressionNode {
	list := &ast.List{Token: parser.currentToken}

	list.Elements = parser.parseExpressionList(token.RIGHTBRACKET)

	return list
}
