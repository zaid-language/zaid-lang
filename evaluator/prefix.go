package evaluator

import (
	"github.com/zaid-language/zaid-lang/ast"
	"github.com/zaid-language/zaid-lang/object"
	"github.com/zaid-language/zaid-lang/value"
)

func evaluatePrefix(node *ast.Prefix, scope *object.Scope) object.Object {
	right := Evaluate(node.Right, scope)

	if isError(right) {
		return right
	}

	switch node.Operator {
	case "!":
		switch right {
		case value.TRUE:
			return value.FALSE
		case value.FALSE:
			return value.TRUE
		case value.NULL:
			return value.TRUE
		default:
			return value.FALSE
		}
	case "-":
		// Only works with number objects
		if right.Type() != object.NUMBER {
			return newError("%d:%d:%s: runtime error: unknown operator: -%s", node.Token.Line, node.Token.Column, node.Token.File, right.Type())
		}

		numberValue := right.(*object.Number).Value.Neg()

		return &object.Number{Value: numberValue}
	}

	return newError("%d:%d:%s: runtime error: unknown operator: %s%s", node.Token.Line, node.Token.Column, node.Token.File, node.Operator, right.Type())
}
