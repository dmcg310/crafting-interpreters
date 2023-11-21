package ast

import "github.com/dmcg310/glox/src/token"

type Expr interface {
	Accept(visitor Visitor) (interface{}, error)
}
type Assign struct {
	Name  token.Token
	Value Expr
}

func (expr *Assign) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitAssign(expr)
}

type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (expr *Binary) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitBinary(expr)
}

type Grouping struct {
	Expression Expr
}

func (expr *Grouping) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitGrouping(expr)
}

type Literal struct {
	Value interface{}
}

func (expr *Literal) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitLiteral(expr)
}

type Logical struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (expr *Logical) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitLogical(expr)
}

type Unary struct {
	Operator token.Token
	Right    Expr
}

func (expr *Unary) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitUnary(expr)
}

type Variable struct {
	Name token.Token
}

func (expr *Variable) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitVariable(expr)
}
