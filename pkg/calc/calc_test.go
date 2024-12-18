package calc

import (
	"testing"

	"github.com/bulbosaur/web-calculator-golang/internal/models"
)

func TestCalc(t *testing.T) {
	cases := []struct {
		expression string
		wantValue  float64
		wantError  error
	}{
		{
			expression: "2+2",
			wantValue:  4,
			wantError:  nil,
		},
		{
			expression: "3*25",
			wantValue:  75,
			wantError:  nil,
		},
		{
			expression: "8/2",
			wantValue:  4,
			wantError:  nil,
		},
		{
			expression: "55 - 3",
			wantValue:  52,
			wantError:  nil,
		},
		{
			expression: "(2 + 2) * 2",
			wantValue:  8,
			wantError:  nil,
		},
		{
			expression: "",
			wantValue:  0,
			wantError:  models.ErrorEmptyExpression,
		},
		{
			expression: "5/0",
			wantValue:  0,
			wantError:  models.ErrorDivisionByZero,
		},
		{
			expression: "2+apple",
			wantValue:  0,
			wantError:  models.ErrorInvalidInput,
		},
		{
			expression: "2 + 2 * 3",
			wantValue:  8,
			wantError:  nil,
		},
		{
			expression: "32 5",
			wantValue:  0,
			wantError:  models.ErrorMissingOperand,
		},
		{
			expression: "((2+3)*2)",
			wantValue:  10,
			wantError:  nil,
		},
		{
			expression: "2++2",
			wantValue:  0,
			wantError:  models.ErrorInvalidInput,
		},
	}

	for _, tc := range cases {
		t.Run(tc.expression, func(t *testing.T) {
			got, err := Calc(tc.expression)
			if err != tc.wantError {
				t.Errorf("Calc(%v) error = %v, wantErr %v", tc.expression, err, tc.wantError)
				return
			}

			if tc.wantError != nil {
				return
			}

			if got != tc.wantValue {
				t.Errorf("Calc(%v) = %v, want %v", tc.expression, got, tc.wantValue)
			}
		})
	}
}
