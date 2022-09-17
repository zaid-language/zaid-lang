package evaluator

import (
	"github.com/zaid-language/zaid-lang/ast"
	"github.com/zaid-language/zaid-lang/object"
	"github.com/zaid-language/zaid-lang/value"
)

func evaluateContinue(node *ast.Continue, scope *object.Scope) object.Object {
	return value.CONTINUE
}
