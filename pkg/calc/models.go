package calc

import "errors"

var ErrorEmptyExpression = errors.New("expression is empty")
var ErrorInvalidInput = errors.New("expression is not valid")
var ErrorDevisionByZero = errors.New("division by zero is not allowed")

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
