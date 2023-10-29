package evaluator

import (
	"zaidlang.org/x/zaid/ast"
	"zaidlang.org/x/zaid/object"
	"zaidlang.org/x/zaid/value"
)

func evaluateNull(node *ast.Null, scope *object.Scope) object.Object {
	return value.NULL
}
