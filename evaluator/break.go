package evaluator

import (
	"zaidlang.org/x/zaid/ast"
	"zaidlang.org/x/zaid/object"
	"zaidlang.org/x/zaid/value"
)

func evaluateBreak(node *ast.Break, scope *object.Scope) object.Object {
	return value.BREAK
}
