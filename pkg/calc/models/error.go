package models

import "errors"

var (
	ErrorDivisionByZero  = errors.New("division by zero is not allowed")
	ErrorEmptyExpression = errors.New("expression is empty")
	ErrorInvalidInput    = errors.New("expression is not valid")
	ErrorInvalidOperand  = errors.New("an invalid operand")
	ErrorMissingOperand  = errors.New("missing operand")
	ErrorUnclosedBracket = errors.New("the brackets in the expression are not consistent")
)
