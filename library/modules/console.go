package modules

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"zaidlang.tech/x/zaid/object"
	"zaidlang.tech/x/zaid/token"
	"zaidlang.tech/x/zaid/value"
)

var ConsoleMethods = map[string]*object.LibraryFunction{}
var ConsoleProperties = map[string]*object.LibraryProperty{}

func init() {
	RegisterMethod(ConsoleMethods, "error", consoleError)
	RegisterMethod(ConsoleMethods, "info", consoleInfo)
	RegisterMethod(ConsoleMethods, "log", consoleLog)
	RegisterMethod(ConsoleMethods, "read", consoleRead)
	RegisterMethod(ConsoleMethods, "warn", consoleWarn)
}

func consoleError(scope *object.Scope, tok token.Token, args ...object.Object) object.Object {
	values := make([]string, 0)

	for _, value := range args {
		values = append(values, value.String())
	}

	printftw(values, "error")

	return nil
}

func consoleInfo(scope *object.Scope, tok token.Token, args ...object.Object) object.Object {
	values := make([]string, 0)

	for _, value := range args {
		values = append(values, value.String())
	}

	printftw(values, "info")

	return nil
}

func consoleLog(scope *object.Scope, tok token.Token, args ...object.Object) object.Object {
	values := make([]string, 0)

	for _, value := range args {
		values = append(values, value.String())
	}

	printftw(values, "")

	return nil
}

func consoleRead(scope *object.Scope, tok token.Token, args ...object.Object) object.Object {
	scanner := bufio.NewScanner(os.Stdin)

	if len(args) == 1 {
		prompt := args[0].(*object.String).Value

		fmt.Print(prompt)
	}

	val := scanner.Scan()

	if !val {
		return value.NULL
	}

	return &object.String{Value: scanner.Text()}
}

func consoleWarn(scope *object.Scope, tok token.Token, args ...object.Object) object.Object {
	values := make([]string, 0)

	for _, value := range args {
		values = append(values, value.String())
	}

	printftw(values, "warning")

	return nil
}

//

func printftw(values []string, prefix string) {
	if len(values) > 0 {
		str := make([]string, 0)

		if len(prefix) > 0 {
			str = append(str, prefix+":")
		}

		str = append(str, values...)

		fmt.Println(strings.Join(str, " "))
	} else {
		fmt.Println()
	}
}