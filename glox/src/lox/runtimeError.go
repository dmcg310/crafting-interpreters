package lox

import (
	"fmt"
	"github.com/dmcg310/glox/src/token"
)

type RuntimeError struct {
	Token token.Token
	Msg   string
}

func (e *RuntimeError) Error() string {
	return fmt.Sprintf("%s: %s", e.Token.Lexeme, e.Msg)
}
