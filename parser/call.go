package parser

import (
	"zaidlang.tech/x/zaid/ast"
	"zaidlang.tech/x/zaid/token"
)

func (parser *Parser) callExpression(callee ast.ExpressionNode) ast.ExpressionNode {
	call := &ast.Call{Token: parser.currentToken, Callee: callee}

	call.Arguments = parser.parseExpressionList(token.RIGHTPAREN)

	return call
}
