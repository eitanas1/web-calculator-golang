package calc

import "errors"

var ErrorEmptyExpression = errors.New("expression is empty")
var ErrorInvalidInput = errors.New("expression is not valid")
var ErrorDevisionByZero = errors.New("division by zero is not allowed")

type Token struct {
	Value    string
	IsNumber bool
}
