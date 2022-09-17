package parser

import (
	"github.com/zaid-language/zaid/ast"
	"github.com/zaid-language/zaid/token"
)

func (parser *Parser) statement() ast.StatementNode {
	switch parser.currentToken.Type {
	case token.RETURN:
		return parser.returnStatement()
	}

	statement := parser.assign()

	if statement != nil {
		return statement
	}

	return parser.expressionStatement()
}

func (parser *Parser) expressionStatement() ast.StatementNode {
	statement := &ast.Expression{}
	statement.Expression = parser.parseExpression(LOWEST)

	return statement
}
