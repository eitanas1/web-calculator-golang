package models

type Token struct {
	Value    string
	IsNumber bool
}

func NewToken(value string, isNumber bool) *Token {
	new_token := Token{
		Value:    value,
		IsNumber: isNumber,
	}
	return &new_token
}
