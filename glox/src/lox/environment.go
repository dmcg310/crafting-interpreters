package lox

import "github.com/dmcg310/glox/src/token"

type Environment struct {
	values    map[string]interface{}
	enclosing *Environment
}

func NewEnvironment(enclosing ...*Environment) Environment {
	env := Environment{
		values: make(map[string]interface{}),
	}

	// function overloading :(
	if len(enclosing) > 0 && enclosing[0] != nil {
		env.enclosing = enclosing[0]
	}

	return env
}

func (e *Environment) get(name token.Token) (interface{}, *RuntimeError) {
	if val, ok := e.values[name.Lexeme]; ok {
		return val, nil
	}

	if e.enclosing != nil {
		return e.enclosing.get(name)
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

	if e.enclosing != nil {
		err := e.enclosing.assign(name, newVal)
		if err != nil {
			return err
		}

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
