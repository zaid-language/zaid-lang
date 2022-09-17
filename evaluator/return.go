package evaluator

import (
	"github.com/zaid-language/zaid-lang/ast"
	"github.com/zaid-language/zaid-lang/object"
)

func evaluateReturn(node *ast.Return, scope *object.Scope) object.Object {
	value := Evaluate(node.Value, scope)

	return &object.Return{Value: value}
}
