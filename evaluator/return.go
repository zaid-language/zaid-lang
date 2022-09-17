package evaluator

import (
	"zaidlang.tech/x/zaid/ast"
	"zaidlang.tech/x/zaid/object"
)

func evaluateReturn(node *ast.Return, scope *object.Scope) object.Object {
	value := Evaluate(node.Value, scope)

	return &object.Return{Value: value}
}
