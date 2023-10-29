package evaluator

import (
	"zaidlang.org/x/zaid/ast"
	"zaidlang.org/x/zaid/object"
)

func evaluateReturn(node *ast.Return, scope *object.Scope) object.Object {
	value := Evaluate(node.Value, scope)

	return &object.Return{Value: value}
}
