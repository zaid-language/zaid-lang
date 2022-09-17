package evaluator

import (
	"zaidlang.tech/x/zaid/ast"
	"zaidlang.tech/x/zaid/object"
	"zaidlang.tech/x/zaid/value"
)

func evaluateContinue(node *ast.Continue, scope *object.Scope) object.Object {
	return value.CONTINUE
}
