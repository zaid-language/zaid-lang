package evaluator

import (
	"github.com/zaid-language/zaid/ast"
	"github.com/zaid-language/zaid/library"
	"github.com/zaid-language/zaid/object"
)

func evaluateIdentifier(node *ast.Identifier, scope *object.Scope) object.Object {
	if libraryModule, ok := library.Modules[node.Value]; ok {
		return libraryModule
	}

	if libraryFunction, ok := library.Functions[node.Value]; ok {
		return libraryFunction
	}

	if identifier, ok := scope.Environment.Get(node.Value); ok {
		return identifier
	}

	return newError("%d:%d:%s: runtime error: unknown identifier: %s", node.Token.Line, node.Token.Column, node.Token.File, node.Value)
}
