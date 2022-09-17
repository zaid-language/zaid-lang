package zaid

import (
	"zaidlang.tech/x/zaid/evaluator"
	"zaidlang.tech/x/zaid/library"
	"zaidlang.tech/x/zaid/library/modules"
	"zaidlang.tech/x/zaid/log"
	"zaidlang.tech/x/zaid/object"
	"zaidlang.tech/x/zaid/parser"
	"zaidlang.tech/x/zaid/scanner"
	"zaidlang.tech/x/zaid/value"
	"zaidlang.tech/x/zaid/version"
)

type Zaid struct {
	FatalError bool
	source     string
	file       string
	Scope      *object.Scope
}

var (
	// Version represents the current version.
	Version = version.Version

	// NULL represents a null value.
	NULL = value.NULL

	// TRUE represents a true value.
	TRUE = value.TRUE

	// FALSE represents a false value.
	FALSE = value.FALSE
)

func New() *Zaid {
	scope := &object.Scope{
		Environment: object.NewEnvironment(),
	}

	zaid := &Zaid{
		Scope: scope,
	}

	zaid.registerEvaluator()

	return zaid
}

func (zaid *Zaid) SetDirectory(directory string) {
	zaid.Scope.Environment.SetDirectory(directory)
}

func (zaid *Zaid) GetDirectory() string {
	return zaid.Scope.Environment.GetDirectory()
}

func (zaid *Zaid) SetSource(source string) {
	zaid.source = source
}

func (zaid *Zaid) SetFile(file string) {
	zaid.file = file
}

func (zaid *Zaid) Execute() object.Object {
	scanner := scanner.New(zaid.source, zaid.file)
	parser := parser.New(scanner)
	program := parser.Parse()

	if len(parser.Errors()) != 0 {
		logParseErrors(parser.Errors())
		return nil
	}

	result := evaluator.Evaluate(program, zaid.Scope)

	if err, ok := result.(*object.Error); ok {
		log.Error(err.Message)

		return nil
	}

	return result
}

func RegisterFunction(name string, function object.GoFunction) {
	library.RegisterFunction(name, function)
}

func RegisterModule(name string, methods map[string]*object.LibraryFunction, properties map[string]*object.LibraryProperty) {
	library.RegisterModule(name, methods, properties)
}

func (zaid *Zaid) registerEvaluator() {
	evaluatorInstance := evaluator.Evaluate

	object.RegisterEvaluator(evaluatorInstance)
	modules.RegisterEvaluator(evaluatorInstance)
}

func logParseErrors(errors []string) {
	for _, message := range errors {
		log.Error(message)
	}
}
