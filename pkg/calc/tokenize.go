package calc

import (
	"unicode"

	"github.com/bulbosaur/web-calculator-golang/internal/models"
)

func tokenize(expression string) ([]models.Token, error) {

	var (
		tokens []models.Token
		number string
		err    error
	)

	for i, symbol := range expression {
		if unicode.IsSpace(symbol) {
			if number != "" {
				tokens = append(tokens, *models.NewToken(number, true))
			}
			continue
		}

		if unicode.IsDigit(symbol) {
			number += string(symbol)
			if i+1 == len(expression) || !unicode.IsDigit(rune(expression[i+1])) {
				tokens = append(tokens, *models.NewToken(number, true))
				number = ""
			}
			continue
		}

		switch string(symbol) {
		case "+", "-", "/", "*", "(", ")":
			tokens = append(tokens, *models.NewToken(string(symbol), false))
		default:
			err = models.ErrorInvalidInput
			return nil, err
		}

	}

	if number != "" {
		tokens = append(tokens, *models.NewToken(number, true))
		number = ""
	}

	if !CheckMissingOperand(tokens) {
		return nil, models.ErrorMissingOperand
	}

	return tokens, nil
}

func CheckMissingOperand(tokens []models.Token) bool {
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
