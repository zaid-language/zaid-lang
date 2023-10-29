package parser

import "zaidlang.org/x/zaid/ast"

func (parser *Parser) postfixExpression() ast.ExpressionNode {
	return &ast.Postfix{
		Token:    parser.previousToken,
		Operator: parser.currentToken.Lexeme,
	}
}
