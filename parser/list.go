package parser

import (
	"github.com/zaid-language/zaid/ast"
	"github.com/zaid-language/zaid/token"
)

func (parser *Parser) listLiteral() ast.ExpressionNode {
	list := &ast.List{Token: parser.currentToken}

	list.Elements = parser.parseExpressionList(token.RIGHTBRACKET)

	return list
}
