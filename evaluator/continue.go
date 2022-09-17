package evaluator

import (
	"github.com/zaid-language/zaid/ast"
	"github.com/zaid-language/zaid/object"
	"github.com/zaid-language/zaid/value"
)

func evaluateContinue(node *ast.Continue, scope *object.Scope) object.Object {
	return value.CONTINUE
}
