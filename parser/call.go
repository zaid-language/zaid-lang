package parser

import (
	"github.com/zaid-language/zaid/ast"
	"github.com/zaid-language/zaid/token"
)

func (parser *Parser) callExpression(callee ast.ExpressionNode) ast.ExpressionNode {
	call := &ast.Call{Token: parser.currentToken, Callee: callee}

	call.Arguments = parser.parseExpressionList(token.RIGHTPAREN)

	return call
}
