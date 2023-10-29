package token

import (
	"fmt"
)

type Token struct {
	Type    TTokentype
	Lexeme  string
	Literal interface{}
	Line    int
}

func NewToken(t TTokentype, lexeme string, literal interface{}, line int) Token {
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
