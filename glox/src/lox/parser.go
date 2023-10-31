package lox

import (
	"github.com/dmcg310/glox/src/ast"
	"github.com/dmcg310/glox/src/token"
)

type Parser struct {
	Tokens  []token.Token
	Current int
}

func NewParser(tokens []token.Token) Parser {
	return Parser{
		Tokens:  tokens,
		Current: 0,
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
		expr = ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr
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
