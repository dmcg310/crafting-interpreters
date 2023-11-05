package lox

import (
	"fmt"
	"github.com/dmcg310/glox/src/ast"
	"github.com/dmcg310/glox/src/token"
)

type Parser struct {
	Tokens  []token.Token
	Current int
	Lox     *Lox
}

func NewParser(tokens []token.Token, lox *Lox) Parser {
	return Parser{
		Tokens:  tokens,
		Current: 0,
		Lox:     lox,
	}
}

func (p *Parser) Parse() []ast.Stmt {
	statements := []ast.Stmt{}
	for !p.isAtEnd() {
		res, _ := p.statement()
		statements = append(statements, res)
	}

	return statements
}

func (p *Parser) expression() (ast.Expr, error) {
	return p.equality()
}

func (p *Parser) statement() (ast.Stmt, error) {
	if p.match(token.PRINT) {
		return p.printStatement()
	}

	return p.expressionStatement()
}

func (p *Parser) printStatement() (ast.Stmt, error) {
	value, err := p.expression()
	if err != nil {
		return nil, err
	}

	p.consume(token.SEMICOLON, "Expect ';' after value.")

	return ast.Stmt.Print(value)
}

func (p *Parser) expressionStatement() (ast.Stmt, error) {
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}

	p.consume(token.SEMICOLON, "Expect ';' after expression.")

	return ast.Stmt.Expression(expr)
}

func (p *Parser) equality() (ast.Expr, error) {
	expr, err := p.comparison()
	if err != nil {
		return nil, err
	}

	for p.match(token.BANG_EQUAL, token.EQUAL_EQUAL) {
		operator := p.previous()
		right, err := p.comparison()
		if err != nil {
			return nil, err
		}

		expr = &ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr, nil
}

func (p *Parser) comparison() (ast.Expr, error) {
	expr, err := p.term()
	if err != nil {
		return nil, err
	}

	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}

		expr = &ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr, nil
}

func (p *Parser) term() (ast.Expr, error) {
	expr, err := p.factor()
	if err != nil {
		return nil, err
	}

	for p.match(token.MINUS, token.PLUS) {
		operator := p.previous()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}

		expr = &ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr, nil
}

func (p *Parser) factor() (ast.Expr, error) {
	expr, err := p.unary()
	if err != nil {
		return nil, err
	}

	for p.match(token.SLASH, token.STAR) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}

		expr = &ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr, nil
}

func (p *Parser) unary() (ast.Expr, error) {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}

		return &ast.Unary{
			Operator: operator,
			Right:    right,
		}, nil
	}

	return p.primary()
}

func (p *Parser) primary() (ast.Expr, error) {
	if p.match(token.FALSE) {
		return &ast.Literal{Value: false}, nil
	}

	if p.match(token.TRUE) {
		return &ast.Literal{Value: true}, nil
	}

	if p.match(token.NIL) {
		return &ast.Literal{Value: nil}, nil
	}

	if p.match(token.NUMBER, token.STRING) {
		return &ast.Literal{Value: p.previous().Literal}, nil
	}

	if p.match(token.LEFT_PAREN) {
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}

		if _, err := p.consume(token.RIGHT_PAREN, "Expect ')' after expression."); err != nil {
			return nil, err
		}

		return &ast.Grouping{Expression: expr}, nil
	}

	return nil, p.error(p.peek(), " Expect expression.")
}

func (p *Parser) match(types ...token.TTokentype) bool {
	for _, ttype := range types {
		if p.check(ttype) {
			p.advance()

			return true
		}
	}

	return false
}

func (p *Parser) consume(ttype token.TTokentype, message string) (token.Token, error) {
	if p.check(ttype) {
		return p.advance(), nil
	}

	return token.Token{}, fmt.Errorf("%v %s", p.peek(), message)
}

func (p *Parser) check(ttype token.TTokentype) bool {
	if p.isAtEnd() {
		return false
	}

	return p.peek().Type == ttype
}

func (p *Parser) advance() token.Token {
	if !p.isAtEnd() {
		p.Current++
	}

	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == token.EOF
}

func (p *Parser) peek() token.Token {
	return p.Tokens[p.Current]
}

func (p *Parser) previous() token.Token {
	return p.Tokens[p.Current-1]
}

func (p *Parser) error(ttoken token.Token, message string) error {
	if ttoken.Type == token.EOF {
		p.Lox.Error(ttoken.Line, " at end"+message)
	} else {
		p.Lox.Error(ttoken.Line, fmt.Sprintf(" at '%s'%s", ttoken.Lexeme, message))
	}

	return NewParserError(ttoken.Line, message)
}

func (p *Parser) synchronise() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().Type == token.SEMICOLON {
			return
		}

		switch p.peek().Type {
		case token.CLASS:
		case token.FUN:
		case token.VAR:
		case token.FOR:
		case token.IF:
		case token.WHILE:
		case token.PRINT:
		case token.RETURN:
			return
		}
	}

	p.advance()
}
