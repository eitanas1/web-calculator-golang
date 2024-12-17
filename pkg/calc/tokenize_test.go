package calc

import (
	"testing"
)

func equalTokens(a, b []Token) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestTokenize(t *testing.T) {
	cases := []struct {
		expression string
		wantOutput []Token
		wantError  error
	}{
		{
			expression: "2+2",
			wantOutput: []Token{{"2", true}, {"+", false}, {"2", true}},
			wantError:  nil,
		},
		{
			expression: "12 + 34",
			wantOutput: []Token{{"12", true}, {"+", false}, {"34", true}},
			wantError:  nil,
		},
		{
			expression: "(5-3)*2",
			wantOutput: []Token{
				{"(", false},
				{"5", true},
				{"-", false},
				{"3", true},
				{")", false},
				{"*", false},
				{"2", true},
			},
			wantError: nil,
		},
		{
			expression: "42",
			wantOutput: []Token{{"42", true}},
			wantError:  nil,
		},
		{
			expression: "2 * (3 + 4)",
			wantOutput: []Token{
				{"2", true},
				{"*", false},
				{"(", false},
				{"3", true},
				{"+", false},
				{"4", true},
				{")", false},
			},
			wantError: nil,
		},
		{
			expression: "",
			wantOutput: []Token{},
			wantError:  nil,
		},
		{
			expression: "2 @ 3",
			wantOutput: nil,
			wantError:  ErrorInvalidInput,
		},
		{
			expression: "2.5 + 3",
			wantOutput: nil,
			wantError:  ErrorInvalidInput,
		},
		{
			expression: "a + b",
			wantOutput: nil,
			wantError:  ErrorInvalidInput,
		},
	}
	for _, tc := range cases {
		t.Run(tc.expression, func(t *testing.T) {
			got, err := tokenize(tc.expression)
			if err != tc.wantError {
				t.Errorf("Tokenize(%v) error = %v, wantErr %v", tc.expression, err, tc.wantError)
				return
			}

			if tc.wantError != nil {
				return
			}

			if !equalTokens(got, tc.wantOutput) {
				t.Errorf("Calc(%v) = %v, want %v", tc.expression, got, tc.wantOutput)
			}
		})
	}
}
