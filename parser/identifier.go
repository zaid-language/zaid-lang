package parser

import (
	"github.com/zaid-language/zaid/ast"
)

func (parser *Parser) identifierLiteral() ast.ExpressionNode {
	identifier := &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Lexeme}

	return identifier
}
