package main

import "fmt"

type Token struct {
	Type    _TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

func NewToken(t _TokenType, lexeme string, literal interface{}, line int) Token {
	return Token{
		Type:    t,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    line,
	}
}

func (t *Token) String() string {
	return fmt.Sprintf("%v %s %v", t.Type, t.Lexeme, t.Literal)
}
