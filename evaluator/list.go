package evaluator

import (
	"zaidlang.tech/x/zaid/ast"
	"zaidlang.tech/x/zaid/object"
)

func evaluateList(node *ast.List, scope *object.Scope) object.Object {
	elements := evaluateExpressions(node.Elements, scope)

	if len(elements) == 1 && isError(elements[0]) {
		return elements[0]
	}

	return &object.List{Elements: elements}
}
