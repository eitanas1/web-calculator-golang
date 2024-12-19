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
			err = models.ErrorInvalidCharacter
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

	if !CheckEmptyBrackets(tokens) {
		return nil, models.ErrorEmptyBrackets
	}

	result := addMissingOperand(tokens)
	return result, nil
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

func CheckEmptyBrackets(tokens []models.Token) bool {
	for i, token := range tokens {
		if i == len(tokens)-1 {
			break
		}
		if token.Value == "(" && tokens[i+1].Value == ")" {
			return false
		}
	}
	return true
}

func addMissingOperand(expression []models.Token) []models.Token {
	var result []models.Token

	for i, token := range expression {
		result = append(result, token)

		if i+1 < len(expression) {
			if token.IsNumber && expression[i+1].Value == "(" {
				result = append(result, models.Token{Value: "*", IsNumber: false})
			}
			if token.Value == ")" && expression[i+1].IsNumber {
				result = append(result, models.Token{Value: "*", IsNumber: false})
			}
		}
	}

	return result
}
