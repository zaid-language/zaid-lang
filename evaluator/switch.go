package evaluator

import (
	"zaidlang.tech/x/zaid/ast"
	"zaidlang.tech/x/zaid/object"
)

func evaluateSwitch(node *ast.Switch, scope *object.Scope) object.Object {
	// Get the value
	obj := Evaluate(node.Value, scope)

	for _, option := range node.Cases {
		for _, val := range option.Value {
			out := Evaluate(val, scope)

			if obj.Type() == out.Type() && (obj.String() == out.String()) {
				// evaluate the block and return the value
				out := evaluateBlock(option.Body, scope)

				return out
			}
		}
	}

	return nil
}
