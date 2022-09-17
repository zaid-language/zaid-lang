package main

import (
	"bytes"
	"syscall/js"

	"zaidlang.tech/x/zaid/evaluator"
	"zaidlang.tech/x/zaid/object"
	"zaidlang.tech/x/zaid/parser"
	"zaidlang.tech/x/zaid/scanner"
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
