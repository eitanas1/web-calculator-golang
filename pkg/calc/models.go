package calc

import "errors"

var ErrorDivisionByZero = errors.New("division by zero is not allowed")
var ErrorEmptyExpression = errors.New("expression is empty")
var ErrorInvalidInput = errors.New("expression is not valid")
var ErrorInvalidOperand = errors.New("an invalid operand")
var ErrorMissingOperand = errors.New("missing operand")
var ErrorUnclosedBracket = errors.New("the brackets in the expression are not consistent")

type Token struct {
	Value    string
	IsNumber bool
}

func newToken(value string, isNumber bool) *Token {
	new_token := Token{
		Value:    value,
		IsNumber: isNumber,
	}
	return &new_token
}

func lastToken(tokens []Token) Token {
	return tokens[len(tokens)-1]
}
