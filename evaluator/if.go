package evaluator

import (
	"github.com/zaid-language/zaid/ast"
	"github.com/zaid-language/zaid/object"
	"github.com/zaid-language/zaid/value"
)

func evaluateIf(node *ast.If, scope *object.Scope) object.Object {
	condition := Evaluate(node.Condition, scope)

	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Evaluate(node.Consequence, scope)
	} else if node.Alternative != nil {
		return Evaluate(node.Alternative, scope)
	}

	return value.NULL
}
