package evaluator

import (
	"zaidlang.org/x/zaid/ast"
	"zaidlang.org/x/zaid/object"
)

func evaluateMethod(node *ast.Method, scope *object.Scope) object.Object {
	left := Evaluate(node.Left, scope)

	if isError(left) {
		return left
	}

	arguments := evaluateExpressions(node.Arguments, scope)

	if len(arguments) == 1 && isError(arguments[0]) {
		return arguments[0]
	}

	result, _ := left.Method(node.Method.(*ast.Identifier).Value, arguments)

	if isError(result) {
		return result
	}

	switch receiver := left.(type) {
	case *object.Map:
		method := node.Method.(*ast.Identifier)

		property := &object.String{Value: method.Value}

		if function, ok := receiver.Pairs[property.MapKey()]; ok {
			return unwrapCall(node.Token, function.Value, arguments, scope)
		}

		return newError("%d:%d:%s: runtime error: unknown method: %s.%s", node.Token.Line, node.Token.Column, node.Token.File, receiver.Type(), method.Value)
	case *object.Instance:
		method := node.Method.(*ast.Identifier)
		evaluated := evaluateInstanceMethod(node, receiver, method.Value, arguments)

		if isError(evaluated) {
			return evaluated
		}

		return unwrapReturn(evaluated)
	case *object.LibraryModule:
		method := node.Method.(*ast.Identifier)
		module := left.(*object.LibraryModule)

		if function, ok := module.Methods[method.Value]; ok {
			return unwrapCall(node.Token, function, arguments, scope)
		}
	}

	return result
}

func evaluateInstanceMethod(node *ast.Method, receiver *object.Instance, name string, arguments []object.Object) object.Object {
	class := receiver.Class
	method, ok := receiver.Class.Environment.Get(name)

	// If we dont have a method, loop through the super classes and check them.
	// Then check the traits.
	if !ok {
		for class != nil {
			method, ok = class.Environment.Get(name)

			if !ok {
				class = class.Super
			} else {
				class = nil
			}
		}
	}

	// if we dont have a method, check for a trait
	if method == nil {
		for _, trait := range receiver.Class.Traits {
			method, ok = trait.Environment.Get(name)

			if !ok {
				return object.NewError("%d:%d:%s: runtime error: undefined method %s for class %s", node.Token.Line, node.Token.Column, node.Token.File, name, receiver.Class.Name.Value)
			}
		}
	}

	// if we still dont have a method, return an error
	if method == nil {
		return object.NewError("%d:%d:%s: runtime error: undefined method %s for class %s", node.Token.Line, node.Token.Column, node.Token.File, name, receiver.Class.Name.Value)
	}

	switch method := method.(type) {
	case *object.Function:
		env := createFunctionEnvironment(method, arguments)
		scope := &object.Scope{Self: receiver, Environment: env}

		return Evaluate(method.Body, scope)
	default:
		return object.NewError("%d:%d:%s: runtime error: invalid type %T in class %s", node.Token.Line, node.Token.Column, node.Token.File, method, receiver.Class.Name.Value)
	}
}
