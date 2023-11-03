package lox

import (
	"bufio"
	"fmt"
	"github.com/dmcg310/glox/src/report"
	"github.com/dmcg310/glox/src/scanner"
	"log"
	"os"
	"strings"
)

type Lox struct {
	HadError bool
	Reporter report.Reporter
}

func (l *Lox) RunFile(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	l.Run(string(bytes))
	if l.HadError {
		os.Exit(65)
	}
}

func (l *Lox) RunPrompt() {
	_scanner := bufio.NewScanner(os.Stdin)

	for {
		l.HadError = false
		fmt.Print("> ")
		if !_scanner.Scan() {
			break
		}

		line := _scanner.Text()
		line = strings.TrimSuffix(line, "\n")
		l.Run(line)
	}

	if err := _scanner.Err(); err != nil {
		_, e := fmt.Fprintln(os.Stderr, "reading standard input:", err)
		if e != nil {
			log.Fatalf("Error printing: %s", e)
		}
	}
}

func (l *Lox) Run(source string) {
	_scanner := scanner.NewScanner(source, l.Reporter)
	tokens := _scanner.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
}

func (l *Lox) Error(line int, message string) {
	l.Reporter.Error(line, message)
	l.HadError = true
}
