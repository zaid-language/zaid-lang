package evaluator

import (
	"zaidlang.tech/x/zaid/ast"
	"zaidlang.tech/x/zaid/object"
)

func evaluateTernary(node *ast.Ternary, scope *object.Scope) object.Object {
	condition := Evaluate(node.Condition, scope)

	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Evaluate(node.IfTrue, scope)
	}

	return Evaluate(node.IfFalse, scope)
}
