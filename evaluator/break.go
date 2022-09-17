package evaluator

import (
	"zaidlang.tech/x/zaid/ast"
	"zaidlang.tech/x/zaid/object"
	"zaidlang.tech/x/zaid/value"
)

func evaluateBreak(node *ast.Break, scope *object.Scope) object.Object {
	return value.BREAK
}
