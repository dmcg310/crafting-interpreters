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

func (p *Parser) expression() ast.Expr {
	return p.equality()
}

func (p *Parser) equality() ast.Expr {
	expr := p.comparison()

	for p.match(token.BANG_EQUAL, token.BANG_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = &ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr
}

func (p *Parser) comparison() ast.Expr {
	expr := p.term()

	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = &ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr
}

func (p *Parser) term() ast.Expr {
	expr := p.factor()

	for p.match(token.MINUS, token.PLUS) {
		operator := p.previous()
		right := p.factor()
		expr = &ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr
}

func (p *Parser) factor() ast.Expr {
	expr := p.unary()

	for p.match(token.SLASH, token.STAR) {
		operator := p.previous()
		right := p.unary()
		expr = &ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr
}

func (p *Parser) unary() ast.Expr {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right := p.unary()

		return &ast.Unary{
			Operator: operator,
			Right:    right,
		}
	}

	return p.primary()
}

func (p *Parser) primary() ast.Expr {
	if p.match(token.FALSE) {
		return &ast.Literal{
			Value: false,
		}
	}

	if p.match(token.TRUE) {
		return &ast.Literal{
			Value: true,
		}
	}

	if p.match(token.NIL) {
		return &ast.Literal{
			Value: nil,
		}
	}

	if p.match(token.NUMBER, token.STRING) {
		return &ast.Literal{
			Value: p.previous().Literal,
		}
	}

	if p.match(token.LEFT_PAREN) {
		expr := p.expression()
		_, err := p.consume(token.RIGHT_PAREN, "Expect ')' after expression.")
		if err != nil {
		}

		return &ast.Grouping{
			Expression: expr,
		}
	}

	return nil
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

	return token.Token{}, fmt.Errorf("%s %s", p.peek(), message)
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

func (p *Parser) error(ttoken token.Token, message string) *ParseError {
	if ttoken.Type == token.EOF {
		p.Lox.Error(ttoken.Line, " at end"+message)
	} else {
		p.Lox.Error(ttoken.Line, fmt.Sprintf(" at '%s'%s", ttoken.Lexeme, message))
	}

	return NewParserError(ttoken.Line, message)
}
