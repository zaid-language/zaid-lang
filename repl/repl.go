package repl

import (
	"io"
	"os"
	"path/filepath"

	"github.com/zaid-language/zaid-lang/log"
	"github.com/zaid-language/zaid-lang/zaid"
	"github.com/peterh/liner"
)

var (
	prompt  = ">> "
	history = filepath.Join(os.TempDir(), ".zaid_history")
)

func Start(in io.Reader, out io.Writer) {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)
	line.SetTabCompletionStyle(liner.TabPrints)

	_, err := os.Open(history)
	if err != nil {
		f, err := os.Create(history)
		if err != nil {
			log.Error("system error: unable to write to history file: %s", err)
		} else {
			line.WriteHistory(f)
			f.Close()
		}
	}

	zaid := zaid.New()

	for {
		source, err := line.Prompt(prompt)

		if err == liner.ErrPromptAborted {
			log.Info("Exiting...")
			os.Exit(1)
		} else {
			evaluate(zaid, source)

			line.AppendHistory(source)
		}
	}
}

func evaluate(zaid *zaid.Zaid, source string) {
	directory, _ := os.Getwd()

	zaid.SetSource(source)
	zaid.SetFile("repl.zaid")
	zaid.SetDirectory(directory)

	result := zaid.Execute()

	if result != nil {
		log.Info(result.String())
	}
}
