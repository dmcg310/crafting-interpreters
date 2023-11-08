package lox

import "github.com/dmcg310/glox/src/token"

type Environment struct {
	values map[string]interface{}
}

func NewEnvironment() Environment {
	return Environment{
		values: make(map[string]interface{}),
	}
}

func (e *Environment) get(name token.Token) (interface{}, *RuntimeError) {
	if val, ok := e.values[name.Lexeme]; ok {
		return val, nil
	}

	return nil, &RuntimeError{
		Token: name,
		Msg:   "Undefined variable '" + name.Lexeme + "'.",
	}
}

func (e *Environment) assign(name token.Token, newVal interface{}) *RuntimeError {
	if _, ok := e.values[name.Lexeme]; ok {
		e.values[name.Lexeme] = newVal

		return nil
	}

	return &RuntimeError{
		Token: name,
		Msg:   "Undefined variable '" + name.Lexeme + "'.",
	}
}

func (e *Environment) define(name string, value interface{}) {
	e.values[name] = value
}
