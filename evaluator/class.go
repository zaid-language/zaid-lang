package evaluator

import (
	"github.com/zaid-language/zaid-lang/ast"
	"github.com/zaid-language/zaid-lang/object"
)

func evaluateClass(node *ast.Class, scope *object.Scope) object.Object {
	class := &object.Class{
		Name:        node.Name,
		Scope:       scope,
		Environment: object.NewEnvironment(),
		Super:       nil,
	}

	// super
	if node.Super != nil {
		identifier, ok := scope.Environment.Get(node.Super.Value)

		if !ok {
			object.NewError("%d:%d:%s: runtime error: identifier '%s' not found in '%s'", node.Super.Token.Line, node.Super.Token.Column, node.Super.Token.File, node.Super.Value, scope.Self.String())
		}

		super, ok := identifier.(*object.Class)

		if !ok {
			object.NewError("%d:%d:%s: runtime error: referenced identifier in extends not a class, got=%T", node.Super.Token.Line, node.Super.Token.Column, node.Super.Token.File, super)
		}

		class.Super = super
	}

	// Create a new scope for this class
	classEnvironment := object.NewEnclosedEnvironment(scope.Environment)
	classScope := &object.Scope{Environment: classEnvironment, Self: class}

	Evaluate(node.Body, classScope)

	scope.Environment.Set(node.Name.Value, class)

	return class
}
