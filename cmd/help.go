package main

import "fmt"

func helpCommand() {
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("    zaid [flags] {file}")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println()
	fmt.Println("    -h  show help")
	fmt.Println("    -t  display how long the program ran for")
	fmt.Println("    -v  show version")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println()
	fmt.Println("    zaid")
	fmt.Println()
	fmt.Println("            Start Zaid REPL")
	fmt.Println()
	fmt.Println("    zaid example.zaid")
	fmt.Println()
	fmt.Println("            Execute source file (example.zaid)")
	fmt.Println()
	fmt.Println("    zaid -t example.zaid")
	fmt.Println()
	fmt.Println("            Execute source file (example.zaid)")
	fmt.Println("            and display how long the program ran for")
	fmt.Println()
	fmt.Println()
}
