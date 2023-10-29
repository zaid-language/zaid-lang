package evaluator

import (
	"zaidlang.org/x/zaid/ast"
	"zaidlang.org/x/zaid/object"
	"zaidlang.org/x/zaid/value"
)

func evaluateContinue(node *ast.Continue, scope *object.Scope) object.Object {
	return value.CONTINUE
}
