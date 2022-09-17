package evaluator

import (
	"github.com/zaid-language/zaid-lang/ast"
	"github.com/zaid-language/zaid-lang/object"
	"github.com/zaid-language/zaid-lang/value"
)

func evaluateNull(node *ast.Null, scope *object.Scope) object.Object {
	return value.NULL
}
