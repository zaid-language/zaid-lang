package parser

import (
	"github.com/zaid-language/zaid-lang/ast"
	"github.com/zaid-language/zaid-lang/token"
)

func (parser *Parser) listLiteral() ast.ExpressionNode {
	list := &ast.List{Token: parser.currentToken}

	list.Elements = parser.parseExpressionList(token.RIGHTBRACKET)

	return list
}
