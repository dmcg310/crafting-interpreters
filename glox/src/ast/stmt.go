package ast

type Stmt interface {
	Accept(visitor Visitor) (interface{}, error)
}

type Expression struct {
	Expression Expr
}

func (stmt *Expression) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitExpression(stmt), nil
}

type Print struct {
	Expression Expr
}

func (stmt *Print) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitPrint(stmt), nil
}
