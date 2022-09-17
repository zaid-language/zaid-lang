package modules

import (
	"path"
	"plugin"
	"strings"

	"github.com/zaid-language/zaid-lang/object"
	"github.com/zaid-language/zaid-lang/parser"
	"github.com/zaid-language/zaid-lang/scanner"
	"github.com/zaid-language/zaid-lang/token"
	"github.com/zaid-language/zaid-lang/version"
)

var ZaidMethods = map[string]*object.LibraryFunction{}
var ZaidProperties = map[string]*object.LibraryProperty{}

func init() {
	RegisterMethod(ZaidMethods, "abort", zaidAbort)
	RegisterMethod(ZaidMethods, "execute", zaidExecute)
	RegisterMethod(ZaidMethods, "extend", zaidExtend)
	RegisterMethod(ZaidMethods, "identifiers", zaidIdentifiers)

	RegisterProperty(ZaidProperties, "version", zaidVersion)
}

func zaidAbort(scope *object.Scope, tok token.Token, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("%d:%d: runtime error: zaid.abort() expects 1 argument. got=%d", tok.Line, tok.Column, len(args))
	}

	switch obj := args[0].(type) {
	case *object.Null:
		return nil
	case *object.String:
		return object.NewError(obj.Value)
	}

	return object.NewError("%d:%d: runtime error: zaid.abort() expects the first argument to be of type 'null' or 'string'. got=%s", tok.Line, tok.Column, strings.ToLower(string(args[0].Type())))
}

func zaidExecute(scope *object.Scope, tok token.Token, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("%d:%d: runtime error: zaid.execute() expects 1 argument. got=%d", tok.Line, tok.Column, len(args))
	}

	source, ok := args[0].(*object.String)

	if !ok {
		return object.NewError("%d:%d: runtime error: zaid.execute() expects the first argument to be of type 'string'. got=%s", tok.Line, tok.Column, strings.ToLower(string(args[0].Type())))
	}

	scanner := scanner.New(source.Value, tok.File)
	parser := parser.New(scanner)
	program := parser.Parse()

	return evaluate(program, scope)
}

func zaidIdentifiers(scope *object.Scope, tok token.Token, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.NewError("%d:%d: runtime error: zaid.identifiers() expects 0 arguments. got=%d", tok.Line, tok.Column, len(args))
	}

	identifiers := []object.Object{}

	store := scope.Environment.All()

	for identifier := range store {
		identifiers = append(identifiers, &object.String{Value: identifier})
	}

	return &object.List{Elements: identifiers}
}

func zaidExtend(scope *object.Scope, tok token.Token, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("%d:%d: runtime error: zaid.extend() expects 1 argument. got=%d", tok.Line, tok.Column, len(args))
	}

	basePath, ok := args[0].(*object.String)

	if !ok {
		return object.NewError("%d:%d: runtime error: zaid.extend() expects the first argument to be of type 'string'. got=%s", tok.Line, tok.Column, strings.ToLower(string(args[0].Type())))
	}

	path := path.Clean(scope.Environment.GetDirectory() + "/" + basePath.Value)

	extension, err := plugin.Open(path)

	if err != nil {
		return object.NewError("%d:%d: runtime error: zaid.extend() failed opening plugin: %s", tok.Line, tok.Column, err)
	}

	register, err := extension.Lookup("Register")

	if err != nil {
		return object.NewError("%d:%d: runtime error: plugin '%s' does not contain Register function: %s", tok.Line, tok.Column, path, err)
	}

	register.(func())()

	return nil
}

// Properties

func zaidVersion(scope *object.Scope, tok token.Token) object.Object {
	return &object.String{Value: version.Version}
}
