package main

import (
	"fmt"
	"github.com/dmcg310/glox/src/lox"
	"github.com/dmcg310/glox/src/report"
	"os"
)

func main() {
	reporter := &report.LoxReporter{}
	l := lox.Lox{
		HadError: false,
		Reporter: reporter,
	}

	if len(os.Args) > 2 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		l.RunFile(os.Args[1])
	} else {
		l.RunPrompt()
	}
}
