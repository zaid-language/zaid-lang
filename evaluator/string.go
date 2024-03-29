package evaluator

import (
	"zaidlang.org/x/zaid/ast"
	"zaidlang.org/x/zaid/object"
)

func evaluateString(node *ast.String, scope *object.Scope) object.Object {
	return &object.String{Value: node.Value}
}

func evaluateStringInfix(node *ast.Infix, left object.Object, right object.Object) object.Object {
	leftValue := left.String()
	rightValue := right.String()

	switch node.Operator {
	case "+":
		return &object.String{Value: leftValue + rightValue}
	case "<":
		return &object.Boolean{Value: leftValue < rightValue}
	case "<=":
		return &object.Boolean{Value: leftValue <= rightValue}
	case ">":
		return &object.Boolean{Value: leftValue > rightValue}
	case ">=":
		return &object.Boolean{Value: leftValue >= rightValue}
	case "==":
		return &object.Boolean{Value: leftValue == rightValue}
	case "!=":
		return &object.Boolean{Value: leftValue != rightValue}
	}

	return newError("%d:%d:%s: runtime error: unknown operator: %s %s %s", node.Token.Line, node.Token.Column, node.Token.File, right.Type(), node.Operator, left.Type())
}
