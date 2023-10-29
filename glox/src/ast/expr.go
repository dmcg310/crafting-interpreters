package ast

import "github.com/dmcg310/glox/src/token"

type Expr interface {
	Accept(visitor Visitor) (interface{}, error)
}

type Binary struct {
	left Expr
	operator token.Token
	right Expr
}

func (expr *Binary) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitBinary(expr)
}

type Grouping struct {
	expression Expr
}

func (expr *Grouping) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitGrouping(expr)
}

type Literal struct {
	value interface{}
}

func (expr *Literal) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitLiteral(expr)
}

type Unary struct {
	operator token.Token
	right Expr
}

func (expr *Unary) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitUnary(expr)
}

