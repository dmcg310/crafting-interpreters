package lox

import "github.com/dmcg310/glox/src/token"

type Enviornment struct {
	values map[string]interface{}
}

func NewEnviornment() Enviornment {
	return Enviornment{
		values: make(map[string]interface{}),
	}
}

func (e *Enviornment) get(name token.Token) (interface{}, *RuntimeError) {
	if val, ok := e.values[name.Lexeme]; ok {
		return val, nil
	}

	return nil, &RuntimeError{
		Token: name,
		Msg:   "Undefined variable '" + name.Lexeme + "'.",
	}
}

func (e *Enviornment) define(name string, value interface{}) {
	e.values[name] = value
}
