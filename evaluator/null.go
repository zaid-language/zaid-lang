package evaluator

import (
	"zaidlang.tech/x/zaid/ast"
	"zaidlang.tech/x/zaid/object"
	"zaidlang.tech/x/zaid/value"
)

func evaluateNull(node *ast.Null, scope *object.Scope) object.Object {
	return value.NULL
}
