package evaluator

import (
	"zaidlang.tech/x/zaid/ast"
	"zaidlang.tech/x/zaid/object"
)

func evaluateBlock(node *ast.Block, scope *object.Scope) object.Object {
	var result object.Object

	for _, statement := range node.Statements {
		result = Evaluate(statement, scope)

		if result != nil {
			resultType := result.Type()

			if resultType == object.ERROR || resultType == object.RETURN || resultType == object.CONTINUE || resultType == object.BREAK {
				return result
			}
		}
	}

	return result
}
