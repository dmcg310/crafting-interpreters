package ast

import "github.com/dmcg310/glox/src/token"

type Stmt interface {
	Accept(visitor Visitor) (interface{}, error)
}

type Block struct {
	Statements []Stmt
}

func (stmt *Block) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitBlock(stmt), nil
}

type Expression struct {
	Expression Expr
}

func (stmt *Expression) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitExpression(stmt), nil
}

type If struct {
	Condition  Expr
	ThenBranch Stmt
	ElseBranch Stmt
}

func (stmt *If) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitIf(stmt), nil
}

type Print struct {
	Expression Expr
}

func (stmt *Print) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitPrint(stmt), nil
}

type Var struct {
	Name        token.Token
	Initialiser Expr
}

func (stmt *Var) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitVar(stmt)
}

type While struct {
	Condition Expr
	Body      Stmt
}

func (stmt *While) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitWhile(stmt)
}
