package lox

import "fmt"

type ParseError struct {
	Line    int
	Message string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("[line %d] Error: %s", e.Line, e.Message)
}

func NewParserError(line int, message string) *ParseError {
	return &ParseError{
		Line:    line,
		Message: message,
	}
}
