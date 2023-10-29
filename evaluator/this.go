package evaluator

import (
	"github.com/zaid-language/zaid-lang/ast"
	"github.com/zaid-language/zaid-lang/object"
)
func evaluateThis(node *ast.This, scope *object.Scope) object.Object {
	if scope.Self != nil {
		return scope.Self
	}

	pairs := make(map[object.MapKey]object.MapPair)

	return &object.Map{Pairs: pairs}
}