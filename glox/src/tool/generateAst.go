package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type GenerateAst struct{}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: generate_ast <output directory>")
		os.Exit(64)
	}

	g := GenerateAst{}
	outputDir := os.Args[1]
	exprTypes := []string{
		"Assign   : token.Token Name, Expr Value",
		"Binary   : Expr Left, token.Token Operator, Expr Right",
		"Grouping : Expr Expression",
		"Literal  : interface{} Value",
		"Logical  : Expr Left, token.Token Operator, Expr Right",
		"Unary    : token.Token Operator, Expr Right",
		"Variable : token.Token Name",
	}

	stmtTypes := []string{
		"Statements : []Stmt",
		"Expression : Expr Expression",
		"If         : Condition Expr, ThenBranch Stmt, ElseBranch Stmt",
		"Print      : Expr Expression",
		"Var        : token.Token Name, Initialiser Expr",
	}

	g.defineAst(outputDir, "Expr", exprTypes)
	g.defineAst(outputDir, "Stmt", stmtTypes)
}

func (g *GenerateAst) defineAst(outputDir, baseName string, types []string) {
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating directory: %s\n", err)
		os.Exit(65)
	}

	path := fmt.Sprintf("%s/%s.go", outputDir, strings.ToLower(baseName))
	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)
		os.Exit(65)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString("package ast\n\n")
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		os.Exit(66)
	}

	_, err = writer.WriteString("import \"github.com/dmcg310/glox/src/token\"\n\n")
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		os.Exit(66)
	}

	_, err = writer.WriteString(fmt.Sprintf("type %s interface {\n\tAccept(visitor Visitor) (interface{}, error)\n}\n\n", baseName))
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		os.Exit(66)
	}

	for _, typ := range types {
		split := strings.Split(typ, ":")
		className := strings.TrimSpace(split[0])
		fields := strings.TrimSpace(split[1])
		g.defineType(writer, baseName, className, fields)
	}

	err = writer.Flush()
	if err != nil {
		fmt.Printf("Error flushing writer: %s\n", err)
		os.Exit(67)
	}
}

func (g *GenerateAst) defineType(writer *bufio.Writer, baseName, className, fieldList string) {
	_, err := writer.WriteString(fmt.Sprintf("type %s struct {\n", className))
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		os.Exit(68)
	}

	fields := strings.Split(fieldList, ", ")
	for _, field := range fields {
		fieldParts := strings.Split(strings.TrimSpace(field), " ")
		if len(fieldParts) != 2 {
			fmt.Printf("Error: Field must have a type and a name, got '%s'\n", field)
			os.Exit(68)
		}
		fieldName := fieldParts[1]
		fieldType := fieldParts[0]
		_, err := writer.WriteString(fmt.Sprintf("\t%s %s\n", fieldName, fieldType))
		if err != nil {
			fmt.Printf("Error writing to file: %s\n", err)
			os.Exit(68)
		}
	}

	_, err = writer.WriteString("}\n\n")
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		os.Exit(68)
	}

	_, err = writer.WriteString(fmt.Sprintf("func (expr *%s) Accept(visitor Visitor) (interface{}, error) {\n", className))
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		os.Exit(68)
	}
	_, err = writer.WriteString(fmt.Sprintf("\treturn visitor.Visit%s(expr)\n", className))
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		os.Exit(68)
	}
	_, err = writer.WriteString("}\n\n")
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		os.Exit(68)
	}
}
