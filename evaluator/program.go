package evaluator

import (
	"zaidlang.org/x/zaid/ast"
	"zaidlang.org/x/zaid/object"
)

func evaluateProgram(node *ast.Program, scope *object.Scope) object.Object {
	var result object.Object

	for _, statement := range node.Statements {
		result = Evaluate(statement, scope)

		switch statement := result.(type) {
		case *object.Error:
			return statement
		case *object.Return:
			return statement.Value
		}
	}

	return result
}
