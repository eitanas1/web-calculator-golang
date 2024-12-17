package calc

import (
	"unicode"
)

func newToken(value string, isNumber bool) *Token {
	new_token := Token{
		Value:    value,
		IsNumber: isNumber,
	}
	return &new_token
}

func tokenize(expression string) ([]Token, error) {
	var tokens []Token
	var number string

	for i, symbol := range expression {
		if unicode.IsSpace(symbol) {
			if number != "" {
				tokens = append(tokens, *newToken(number, true))
			}
			continue
		}

		if unicode.IsDigit(symbol) {
			number += string(symbol)
			if i+1 == len(expression) {
				tokens = append(tokens, *newToken(number, true))
				continue
			}
			next_symbol := expression[i+1]
			if !unicode.IsDigit(rune(next_symbol)) {
				tokens = append(tokens, *newToken(number, true))
			}
			continue
		}

		if number != "" {
			tokens = append(tokens, *newToken(number, true))
			number = ""
		}

		switch string(symbol) {
		case "+", "-", "/", "*", "(", ")":
			tokens = append(tokens, *newToken(string(symbol), false))
		default:
			return nil, ErrorInvalidInput
		}
	}

	if number != "" {
		tokens = append(tokens, *newToken(number, true))
		number = ""
	}

	return tokens, nil
}
