package parser

import (
	"zaidlang.org/x/zaid/ast"
	"zaidlang.org/x/zaid/token"
)

func (parser *Parser) callExpression(callee ast.ExpressionNode) ast.ExpressionNode {
	call := &ast.Call{Token: parser.currentToken, Callee: callee}

	call.Arguments = parser.parseExpressionList(token.RIGHTPAREN)

	return call
}
