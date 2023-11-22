package ast

type Visitor interface {
	VisitAssign(expr *Assign) (interface{}, error)
	VisitBinary(expr *Binary) (interface{}, error)
	VisitGrouping(expr *Grouping) (interface{}, error)
	VisitLiteral(expr *Literal) (interface{}, error)
	VisitLogical(expr *Logical) (interface{}, error)
	VisitUnary(expr *Unary) (interface{}, error)
	VisitVariable(expr *Variable) (interface{}, error)

	VisitBlock(expr *Block) error
	VisitExpression(expr *Expression) error
	VisitIf(expr *If) error
	VisitPrint(expr *Print) error
	VisitVar(stmt *Var) (interface{}, error)
	VisitWhile(stmt *While) (interface{}, error)
}
