package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/zaid-language/zaid-lang/log"
	"github.com/zaid-language/zaid-lang/repl"
	"github.com/zaid-language/zaid-lang/version"
	"github.com/zaid-language/zaid-lang/zaid"
)

var (
	flagHelp      bool
	flagVersion   bool
	flagBenchmark bool
	flagTime      bool
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options] [<filename>]\n", path.Base(os.Args[0]))
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.BoolVar(&flagHelp, "h", false, "display help information")
	flag.BoolVar(&flagVersion, "v", false, "display version information")
	flag.BoolVar(&flagBenchmark, "b", false, "run benchmark tests against Zaid and Go")
	flag.BoolVar(&flagTime, "t", false, "display how long the program ran for")
}

func main() {
	flag.Parse()

	if flagVersion {
		fmt.Printf("%s %s\n", path.Base(os.Args[0]), version.Version)
		os.Exit(0)
	}

	if flagHelp {
		helpCommand()
		os.Exit(2)
	}

	if flagBenchmark {
		benchmarkCommand()
		os.Exit(2)
	}

	args := flag.Args()

	if len(args) == 0 {
		fmt.Printf("Zaid (%s)\n", version.Version)
		fmt.Printf("Press Ctrl + C to exit\n\n")

		repl.Start(os.Stdin, os.Stdout)
		return
	}

	if len(args) > 0 {
		start := time.Now()
		sourceFile, err := os.Open(args[0])

		if err != nil {
			log.Error("system error: could not open source file %s: %s", args[0], err)

			os.Exit(1)
		}

		defer sourceFile.Close()

		sourceBuffer := bytes.NewBuffer(nil)
		io.Copy(sourceBuffer, sourceFile)
		source := sourceBuffer.String()

		directory, _ := filepath.Abs(filepath.Dir(args[0]))
		fullPath, _ := filepath.Abs(args[0])
		currentFile := strings.Replace(fullPath, directory+"/", "", 1)

		zaid := zaid.New()
		zaid.SetSource(source)
		zaid.SetFile(currentFile)
		zaid.SetDirectory(directory)
		zaid.Execute()

		elapsed := time.Since(start)

		if flagTime {
			log.Info("(executed in: %s)\n", elapsed)
		}
	}
}
