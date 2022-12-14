package object

import (
	"github.com/zaid-language/zaid-lang/ast"
	"github.com/zaid-language/zaid-lang/token"
)

var evaluator func(node ast.Node, scope *Scope) Object

// Type is the type of the object given as a string.
type Type string

// Object is the interface for all object values.
type Object interface {
	HasMethods
	Type() Type
	String() string
}

type MapKey struct {
	Type  Type
	Value uint64
}

type Mappable interface {
	MapKey() MapKey
}

type HasMethods interface {
	Method(method string, args []Object) (Object, bool)
}

type GoFunction func(scope *Scope, tok token.Token, args ...Object) Object
type GoProperty func(scope *Scope, tok token.Token) Object
type ObjectMethod func(value interface{}, args ...Object) (Object, bool)

func RegisterEvaluator(e func(node ast.Node, scope *Scope) Object) {
	evaluator = e
}
