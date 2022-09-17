package evaluator

import (
	"github.com/zaid-language/zaid/ast"
	"github.com/zaid-language/zaid/object"
	"github.com/zaid-language/zaid/value"
)

func evaluateFor(node *ast.For, scope *object.Scope) object.Object {
	existingIdentifier, identifierExisted := scope.Environment.Get(node.Identifier.Value)

	defer func() {
		if identifierExisted {
			scope.Environment.Set(node.Identifier.Value, existingIdentifier)
		} else {
			scope.Environment.Delete(node.Identifier.Value)
		}
	}()

	initializer := Evaluate(node.Initializer, scope)

	if isError(initializer) {
		return initializer
	}

	loop := true

	for loop {
		condition := Evaluate(node.Condition, scope)

		if isError(condition) {
			return condition
		}

		if isTruthy(condition) {
			err := Evaluate(node.Block, scope)

			if isTerminator(err) {
				switch val := err.(type) {
				case *object.Error:
					return val
				case *object.Continue:
					//
				case *object.Break:
					return nil
				}
			}

			err = Evaluate(node.Increment, scope)

			if isError(err) {
				return err
			}

			continue
		}

		loop = false
	}

	return value.NULL
}
