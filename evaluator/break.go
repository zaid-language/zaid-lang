package evaluator

import (
	"github.com/zaid-language/zaid/ast"
	"github.com/zaid-language/zaid/object"
	"github.com/zaid-language/zaid/value"
)

func evaluateBreak(node *ast.Break, scope *object.Scope) object.Object {
	return value.BREAK
}
