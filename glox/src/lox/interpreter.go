package lox

import (
	"fmt"
	"github.com/dmcg310/glox/src/ast"
	"github.com/dmcg310/glox/src/token"
	"strconv"
)

type Interpreter struct{}

func (i *Interpreter) interpret(statements []ast.Stmt) error {
	for _, stmt := range statements {
		if expressionStmt, ok := stmt.(*ast.Expression); ok {
			result, err := i.evaluate(expressionStmt.Expression)
			if err != nil {
				return err
			}

			fmt.Println(i.stringify(result))
		} else {
			_, err := stmt.Accept(i)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (i *Interpreter) VisitLiteral(expr *ast.Literal) (interface{}, error) {
	return expr.Value, nil
}

func (i *Interpreter) VisitGrouping(expr *ast.Grouping) (interface{}, error) {
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) evaluate(expr ast.Expr) (interface{}, error) {
	return expr.Accept(i)
}

func (i *Interpreter) execute(stmt ast.Stmt) (interface{}, error) {
	return stmt.Accept(i)
}

func (i *Interpreter) VisitExpression(stmt *ast.Expression) error {
	_, err := i.evaluate(stmt.Expression)

	return err
}

func (i *Interpreter) VisitPrint(stmt *ast.Print) error {
	value, err := i.evaluate(stmt.Expression)
	if err != nil {
		return err
	}

	fmt.Println(i.stringify(value))

	return nil
}

// temporary
func (i *Interpreter) VisitVariable(expr *ast.Variable) (interface{}, error) {
	return nil, nil
}

// temporary
func (i *Interpreter) VisitVar(stmt *ast.Var) (interface{}, error) {
	return nil, nil
}

func (i *Interpreter) VisitBinary(expr *ast.Binary) (interface{}, error) {
	left, err := i.evaluate(expr.Left)
	if err != nil {
		return nil, err
	}

	right, err := i.evaluate(expr.Right)
	if err != nil {
		return nil, err
	}

	leftNum, leftIsNum := left.(float64)
	rightNum, rightIsNum := right.(float64)
	if !leftIsNum || !rightIsNum {
		return nil, fmt.Errorf("operands must be numbers")
	}

	switch expr.Operator.Type {
	case token.GREATER:
		err := i.checkNumberOperands(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}

		return leftNum > rightNum, nil
	case token.GREATER_EQUAL:
		err := i.checkNumberOperands(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}

		return leftNum >= rightNum, nil
	case token.LESS:
		err := i.checkNumberOperands(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}

		return leftNum < rightNum, nil
	case token.LESS_EQUAL:
		err := i.checkNumberOperands(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}

		return leftNum <= rightNum, nil
	case token.BANG_EQUAL:
		return !i.isEqual(left, right), nil
	case token.EQUAL_EQUAL:
		return i.isEqual(left, right), nil
	case token.PLUS:
		// handle addition for numbers
		leftNum, leftIsNum := left.(float64)
		rightNum, rightIsNum := right.(float64)
		if leftIsNum && rightIsNum {
			return leftNum + rightNum, nil
		}

		// handle concatenation for strings
		leftStr, leftIsStr := left.(string)
		rightStr, rightIsStr := right.(string)
		if leftIsStr && rightIsStr {
			return leftStr + rightStr, nil
		}

		return nil, &RuntimeError{Token: expr.Operator, Msg: "Operands must be two numbers or two strings."}
	case token.MINUS:
		err := i.checkNumberOperands(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}

		leftVal, ok := left.(float64)
		if !ok {
			return nil, fmt.Errorf("left operand must be a number")
		}

		rightVal, ok := right.(float64)
		if !ok {
			return nil, fmt.Errorf("right operand must be a number")
		}

		return leftVal - rightVal, nil
	case token.SLASH:
		err := i.checkNumberOperands(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}

		leftVal, ok := left.(float64)
		if !ok {
			return nil, fmt.Errorf("left operand must be a number")
		}

		rightVal, ok := right.(float64)
		if !ok {
			return nil, fmt.Errorf("right operand must be a number")
		}

		return leftVal / rightVal, nil
	case token.STAR:
		err := i.checkNumberOperands(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}

		leftVal, ok := left.(float64)
		if !ok {
			return nil, fmt.Errorf("left operand must be a number")
		}

		rightVal, ok := right.(float64)
		if !ok {
			return nil, fmt.Errorf("right operand must be a number")
		}

		return leftVal * rightVal, nil
	}

	return nil, fmt.Errorf("unknown binary operator: %v", expr.Operator.Type)
}

func (i *Interpreter) VisitUnary(expr *ast.Unary) (interface{}, error) {
	right, err := i.evaluate(expr.Right)
	if err != nil {
		return nil, err
	}

	switch expr.Operator.Type {
	case token.MINUS:
		err := i.checkNumberOperand(expr.Operator, right)
		if err != nil {
			return nil, err
		}

		rightVal, ok := right.(float64)
		if !ok {
			return nil, fmt.Errorf("operand must be a number")
		}

		return -rightVal, nil
	case token.BANG:
		return !i.isTruthy(right), nil
	}

	return nil, fmt.Errorf("unknown unary operator: %v", expr.Operator.Type)
}

func (i *Interpreter) checkNumberOperand(operator token.Token, operand interface{}) error {
	_, ok := operand.(float64)
	if !ok {
		return &RuntimeError{Token: operator, Msg: "Operand must be a number."}
	}

	return nil
}

func (i *Interpreter) checkNumberOperands(operator token.Token, left, right interface{}) error {
	_, leftOk := left.(float64)
	_, rightOk := right.(float64)
	if leftOk && rightOk {
		return nil
	}

	return &RuntimeError{Token: operator, Msg: "Operands must be numbers."}
}

func (i *Interpreter) isTruthy(obj interface{}) bool {
	if obj == nil {
		return false
	}

	val, ok := obj.(bool)
	if ok {
		return val
	}

	return true
}

func (i *Interpreter) isEqual(a, b interface{}) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil {
		return false
	}

	return a == b
}

func (i *Interpreter) stringify(obj interface{}) string {
	if obj == nil {
		return "nil"
	}

	if num, ok := obj.(float64); ok {
		return strconv.FormatFloat(num, 'f', -1, 64)
	}

	if str, ok := obj.(string); ok {
		return str
	}

	return fmt.Sprintf("%v", obj)
}
