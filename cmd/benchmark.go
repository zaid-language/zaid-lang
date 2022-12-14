package main

import (
	"fmt"
	"time"

	"github.com/zaid-language/zaid-lang/evaluator"
	"github.com/zaid-language/zaid-lang/object"
	"github.com/zaid-language/zaid-lang/parser"
	"github.com/zaid-language/zaid-lang/scanner"
)

func benchmarkCommand() {
	benchmarkHelloWorld()
}

func benchmarkHelloWorld() {
	goTime := nativeHelloWorld()
	scanTime, parseTime, interpretTime, zaidTime := benchmark(`printftw("Hello, world!")`)

	fmt.Println("==============================")
	fmt.Println("Hello world benchmark")
	fmt.Println("==============================")
	fmt.Printf("Golang:             %s\n", goTime)
	fmt.Printf("github.com/zaid-language/zaid-lang:          %s\n", zaidTime)
	fmt.Printf("-- Scanner:     %s\n", scanTime)
	fmt.Printf("-- Parser:      %s\n", parseTime)
	fmt.Printf("-- Interpreter: %s\n", interpretTime)
}

func nativeHelloWorld() time.Duration {
	start := time.Now()
	fmt.Println("Hello, world!")

	return time.Since(start)
}

func benchmark(source string) (scanTime time.Duration, parseTime time.Duration, interpretTime time.Duration, zaidTime time.Duration) {

	scope := &object.Scope{
		Environment: object.NewEnvironment(),
	}

	start := time.Now()

	scanner := scanner.New(source, "benchmark.zaid")
	scanTime = time.Since(start)

	parseStart := time.Now()
	parser := parser.New(scanner)
	program := parser.Parse()
	parseTime = time.Since(parseStart)

	interpretStart := time.Now()
	evaluator.Evaluate(program, scope)
	interpretTime = time.Since(interpretStart)
	zaidTime = time.Since(start)

	return scanTime, parseTime, interpretTime, zaidTime
}
