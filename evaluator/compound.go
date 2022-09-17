package evaluator

import (
	"github.com/zaid-language/zaid/ast"
	"github.com/zaid-language/zaid/object"
)

func evaluateCompound(node *ast.Compound, scope *object.Scope) object.Object {
	infix := &ast.Infix{
		Token:    node.Token,
		Left:     node.Left,
		Operator: node.Operator[:len(node.Operator)-1],
		Right:    node.Right,
	}

	value := evaluateInfix(infix, scope)

	scope.Environment.Set(node.Left.(*ast.Identifier).Value, value)

	return nil
}
