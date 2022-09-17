package parser

import (
	"github.com/zaid-language/zaid-lang/ast"
	"github.com/zaid-language/zaid-lang/token"
)

func (parser *Parser) callExpression(callee ast.ExpressionNode) ast.ExpressionNode {
	call := &ast.Call{Token: parser.currentToken, Callee: callee}

	call.Arguments = parser.parseExpressionList(token.RIGHTPAREN)

	return call
}
