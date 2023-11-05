package ast

type Visitor interface {
	VisitBinary(expr *Binary) (interface{}, error)
	VisitGrouping(expr *Grouping) (interface{}, error)
	VisitLiteral(expr *Literal) (interface{}, error)
	VisitUnary(expr *Unary) (interface{}, error)
	VisitExpression(expr *Expression) (interface{}, error)
	VisitPrint(expr *Print) (interface{}, error)
}
