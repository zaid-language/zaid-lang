package parser

import (
	"zaidlang.org/x/zaid/ast"
)

func (parser *Parser) identifierLiteral() ast.ExpressionNode {
	identifier := &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Lexeme}

	return identifier
}
