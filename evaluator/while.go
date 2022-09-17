package evaluator

import (
	"github.com/zaid-language/zaid/ast"
	"github.com/zaid-language/zaid/object"
)

func evaluateWhile(node *ast.While, scope *object.Scope) object.Object {
	for {
		condition := Evaluate(node.Condition, scope)

		if isError(condition) {
			return condition
		}

		if isTruthy(condition) {
			evaluated := Evaluate(node.Consequence, scope)

			if isTerminator(evaluated) {
				switch val := evaluated.(type) {
				case *object.Error:
					return val
				case *object.Continue:
					//
				case *object.Break:
					return nil
				}
			}
		} else {
			break
		}
	}

	return nil
}
