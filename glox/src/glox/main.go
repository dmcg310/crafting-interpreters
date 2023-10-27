package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		l.run(line)

		l.HadError = false
	}

	if err := scanner.Err(); err != nil {
		_, e := fmt.Fprintln(os.Stderr, "reading standard input:", err)
		if e != nil {
			log.Fatalf("Error printing: %s", e)
		}
	}
}

func (l *Lox) run(source string) {
	for _, runeValue := range source {
		char := string(runeValue)
		fmt.Println(char)
	}
}

func (l *Lox) error(line int, message string) {
	l.report(line, "", message)
}

func (l *Lox) report(line int, where, message string) {
	log.Printf("[line %d] Error %s: %s\n", line, where, message)
	l.HadError = true
}
