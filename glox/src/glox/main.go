package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		run_file(os.Args[1])
	} else {
		run_prompt()
	}
}

func run_file(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	run(string(bytes))
}

func run_prompt() {
	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		if input == "exit\n" {
			break
		}

		fmt.Print(input)
		run(input)
	}
}

func run(source string) {
	// scanner := NewScanner(source)
	// tokens, err := scanner.ScanTokens()
	// if err != nil {
	// 	fmt.Println("Error scanning tokens:", err)
	// 	return
	// }
	//
	// for _, token := range tokens {
	// 	fmt.Println(token)
	// }
}
