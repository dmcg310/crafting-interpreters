package tool

import (
	"fmt"
	"os"
)

type GenerateAst struct{}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: generate_ast <output directory>")
		os.Exit(64)
	}

	g := GenerateAst{}
	outputDir := os.Args[1]
	types := []string{
		"Binary   : Expr left, Token operator, Expr right",
		"Grouping : Expr expression",
		"Literal  : Object value",
		"Unary    : Token operator, Expr right",
	}
	g.defineAst(outputDir, "Expr", types)
}

func (g *GenerateAst) defineAst(outputDir, baseName string, types []string) {
	path := fmt.Sprintf("%s/%s.go", outputDir, baseName)
	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(65)
	}

	_, _ = file.WriteString("package main")
}
