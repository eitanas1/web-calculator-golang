package calc

import (
	"unicode"
)

func tokenize(expression string) ([]Token, error) {

	var (
		tokens []Token
		number string
		err    error
	)

	for i, symbol := range expression {
		if unicode.IsSpace(symbol) {
			if number != "" {
				tokens = append(tokens, *newToken(number, true))
			}
			continue
		}

		if unicode.IsDigit(symbol) {
			number += string(symbol)
			if i+1 == len(expression) || !unicode.IsDigit(rune(expression[i+1])) {
				tokens = append(tokens, *newToken(number, true))
				number = ""
			}
			continue
		}

		switch string(symbol) {
		case "+", "-", "/", "*", "(", ")":
			tokens = append(tokens, *newToken(string(symbol), false))
		default:
			err = ErrorInvalidInput
			return nil, err
		}

	}

	if number != "" {
		tokens = append(tokens, *newToken(number, true))
		number = ""
	}

	if !CheckMissingOperand(tokens) {
		return nil, ErrorMissingOperand
	}

	return tokens, nil
}

func CheckMissingOperand(tokens []Token) bool {
	for i, token := range tokens {
		if i == len(tokens)-1 {
			break
		}
		if token.IsNumber && tokens[i+1].IsNumber {
			return false
		}
	}
	return true
}
