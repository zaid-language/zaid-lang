package main

import (
	"bytes"
	"syscall/js"

	"github.com/zaid-language/zaid/evaluator"
	"github.com/zaid-language/zaid/object"
	"github.com/zaid-language/zaid/parser"
	"github.com/zaid-language/zaid/scanner"
)

func zaid(this js.Value, i []js.Value) interface{} {
	m := make(map[string]interface{})
	var buf bytes.Buffer

	code := i[0].String()
	scope := &object.Scope{Environment: object.NewEnvironment()}
	scope.Environment.SetWriter(&buf)

	scanner := scanner.New(code, "wasm.zaid")
	parser := parser.New(scanner)
	program := parser.Parse()

	result := evaluator.Evaluate(program, scope)

	m["result"] = buf.String()
	m["object"] = result.String()

	return js.ValueOf(m)
}

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("zaid", js.FuncOf(zaid))
	<-c
}
