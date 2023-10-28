package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Lox struct {
	HadError bool
}

func main() {
	l := Lox{
		HadError: false,
	}

	if len(os.Args) > 2 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		l.runFile(os.Args[1])
	} else {
		l.runPrompt()
	}
}

func (l *Lox) runFile(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	l.run(string(bytes))
	if l.HadError {
		os.Exit(65)
	}
}

func (l *Lox) runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		l.HadError = false
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		line = strings.TrimSuffix(line, "\n")
		l.run(line)
	}

	if err := scanner.Err(); err != nil {
		_, e := fmt.Fprintln(os.Stderr, "reading standard input:", err)
		if e != nil {
			log.Fatalf("Error printing: %s", e)
		}
	}
}

func (l *Lox) run(source string) {
	scanner := _NewScanner(source, l)
	tokens := scanner.scanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
}

func (l *Lox) Error(line int, message string) {
	l.report(line, "", message)
}

func (l *Lox) report(line int, where, message string) {
	log.Printf("[line %d] Error %s: %s\n", line, where, message)
	l.HadError = true
}
