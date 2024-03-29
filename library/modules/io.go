package modules

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"zaidlang.org/x/zaid/object"
	"zaidlang.org/x/zaid/token"
)

var IoMethods = map[string]*object.LibraryFunction{}
var IoProperties = map[string]*object.LibraryProperty{}

func init() {
	RegisterMethod(IoMethods, "append", ioAppend)
	RegisterMethod(IoMethods, "read", ioRead)
	RegisterMethod(IoMethods, "write", ioWrite)
}

func ioAppend(scope *object.Scope, tok token.Token, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.NewError("%d:%d: runtime error: io.append() expects 2 arguments. got=%d", tok.Line, tok.Column, len(args))
	}

	basePath, ok := args[0].(*object.String)

	if !ok {
		return object.NewError("%d:%d: runtime error: io.append() expects first argument to be of type 'string'. got=%s", tok.Line, tok.Column, strings.ToLower(string(args[0].Type())))
	}

	content, ok := args[1].(*object.String)

	if !ok {
		return object.NewError("%d:%d: runtime error: io.append() expects second argument to be of type 'string'. got=%s", tok.Line, tok.Column, strings.ToLower(string(args[1].Type())))
	}

	cleanPath := path.Clean(scope.Environment.GetDirectory() + "/" + basePath.Value)

	file, err := os.OpenFile(cleanPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return object.NewError("%d:%d: runtime error: io.append() %s", tok.Line, tok.Column, err)
	}

	defer file.Close()

	file.WriteString(content.Value + "\n")

	return nil
}

func ioRead(scope *object.Scope, tok token.Token, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("%d:%d: runtime error: io.read() expects 1 argument. got=%d", tok.Line, tok.Column, len(args))
	}

	basePath, ok := args[0].(*object.String)

	if !ok {
		return object.NewError("%d:%d: runtime error: io.read() expects first argument to be of type 'string'. got=%s", tok.Line, tok.Column, strings.ToLower(string(args[0].Type())))
	}

	path := path.Clean(scope.Environment.GetDirectory() + "/" + basePath.Value)
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return object.NewError("%d:%d: runtime error: io.read() %s", tok.Line, tok.Column, err)
	}

	return &object.String{Value: string(content)}
}

func ioWrite(scope *object.Scope, tok token.Token, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.NewError("%d:%d: runtime error: io.write() expects 2 arguments. got=%d", tok.Line, tok.Column, len(args))
	}

	basePath, ok := args[0].(*object.String)

	if !ok {
		return object.NewError("%d:%d: runtime error: io.write() expects first argument to be of type 'string'. got=%s", tok.Line, tok.Column, strings.ToLower(string(args[0].Type())))
	}

	content, ok := args[1].(*object.String)

	if !ok {
		return object.NewError("%d:%d: runtime error: io.write() expects second argument to be of type 'string'. got=%s", tok.Line, tok.Column, strings.ToLower(string(args[1].Type())))
	}

	path := path.Clean(scope.Environment.GetDirectory() + "/" + basePath.Value)
	contents := []byte(content.Value)
	info, err := os.Stat(path)

	if err != nil {
		return object.NewError("%d:%d: runtime error: io.write() %s", tok.Line, tok.Column, err)
	}

	mode := info.Mode()

	err = ioutil.WriteFile(path, contents, mode)

	if err != nil {
		return object.NewError("%d:%d: runtime error: io.write() %s", tok.Line, tok.Column, err)
	}

	return nil
}
