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
	Interpreter     Interpreter
	HadError        bool
	HadRuntimeError bool
	Reporter        report.Reporter
	Environment     Environment
}

func (l *Lox) RunFile(path string) {
	l.Environment = NewEnvironment()

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
	l.Environment = NewEnvironment()
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
	parser := NewParser(tokens, l)
	expr := parser.Parse()

	if l.HadError {
		return
	}

	if l.HadRuntimeError {
		os.Exit(70)
	}

	_ = l.Interpreter.interpret(expr, l.Environment)
}

func (l *Lox) Error(line int, message string) {
	l.Reporter.Error(line, message)
	l.HadError = true
}

func (l *Lox) runtimeError(error *RuntimeError) {
	log.Printf("%s\n[line %d]\n", error.Msg, error.Token.Line)
	l.HadRuntimeError = true
}
