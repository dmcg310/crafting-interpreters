package lox

import (
	"fmt"
	"github.com/dmcg310/glox/src/ast"
	"strings"
)

type AstPrinter struct{}

func (printer *AstPrinter) Print(expr ast.Expr) string {
	result, err := expr.Accept(printer)
	if err != nil {
		fmt.Println("Error during AST printing:", err)
		return ""
	}

	strResult, ok := result.(string)
	if !ok {
		fmt.Println("Accept did not return a string.")

		return ""
	}

	return strResult
}

func (printer *AstPrinter) VisitBinary(expr *ast.Binary) (interface{}, error) {
	return printer.visitBinaryExpr(expr), nil
}

func (printer *AstPrinter) VisitGrouping(expr *ast.Grouping) (interface{}, error) {
	return printer.visitGroupingExpr(expr), nil
}

func (printer *AstPrinter) VisitLiteral(expr *ast.Literal) (interface{}, error) {
	return printer.visitLiteralExpr(expr), nil
}

func (printer *AstPrinter) VisitUnary(expr *ast.Unary) (interface{}, error) {
	return printer.visitUnaryExpr(expr), nil
}

func (printer *AstPrinter) visitBinaryExpr(expr *ast.Binary) string {
	return printer.parenthesise(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (printer *AstPrinter) visitGroupingExpr(expr *ast.Grouping) string {
	return printer.parenthesise("group", expr.Expression)
}

func (printer *AstPrinter) visitLiteralExpr(expr *ast.Literal) string {
	if expr.Value == nil {
		return "nil"
	}

	return fmt.Sprintf("%v", expr.Value)
}

func (printer *AstPrinter) visitUnaryExpr(expr *ast.Unary) string {
	return printer.parenthesise(expr.Operator.Lexeme, expr.Right)
}

func (printer *AstPrinter) parenthesise(name string, exprs ...ast.Expr) string {
	var builder strings.Builder
	builder.WriteString("(")
	builder.WriteString(name)

	for _, expr := range exprs {
		builder.WriteString(" ")

		if str, err := expr.Accept(printer); err == nil {
			builder.WriteString(fmt.Sprint(str))
		} else {
			fmt.Println("Error during Accept:", err)

			return ""
		}
	}
	builder.WriteString(")")

	return builder.String()
}
