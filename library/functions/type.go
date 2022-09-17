package functions

import (
	"strings"

	"zaidlang.tech/x/zaid/object"
	"zaidlang.tech/x/zaid/token"
)

func Type(scope *object.Scope, tok token.Token, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("%d:%d: runtime error: type() expects 1 argument. got=%d", tok.Line, tok.Column, len(args))
	}

	objectType := string(args[0].Type())

	return &object.String{Value: strings.ToLower(objectType)}
}
