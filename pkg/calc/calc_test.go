package calc

import (
	"testing"
)

func TestCalc(t *testing.T) {
	cases := []struct {
		name       string
		expression string
		wantValue  float64
		wantError  error
	}{
		{
			name:       "Simple addition",
			expression: "2+2",
			wantValue:  4,
			wantError:  nil,
		},
		{
			name:       "Simple multiplication",
			expression: "3*3",
			wantValue:  9,
			wantError:  nil,
		},
		{
			name:       "Simple division",
			expression: "8/2",
			wantValue:  4,
			wantError:  nil,
		},
		{
			name:       "Simple subtraction",
			expression: "5-3",
			wantValue:  2,
			wantError:  nil,
		},
		{
			name:       "Complex expression",
			expression: "(2+2)*2",
			wantValue:  8,
			wantError:  nil,
		},
		{
			name:       "Empty expression",
			expression: "",
			wantValue:  0,
			wantError:  ErrorInvalidInput,
		},
		{
			name:       "Division by zero",
			expression: "5/0",
			wantValue:  0,
			wantError:  ErrorDevisionByZero,
		},
		{
			name:       "Invalid characters",
			expression: "2+apple",
			wantValue:  0,
			wantError:  ErrorInvalidInput,
		},
		{
			name:       "Multiple operations",
			expression: "2+2*3",
			wantValue:  8,
			wantError:  nil,
		},
		{
			name:       "Brackets in brackets",
			expression: "((2+3)*2)",
			wantValue:  10,
			wantError:  nil,
		},
		{
			name:       "Invalid expression",
			expression: "2++2",
			wantValue:  0,
			wantError:  ErrorInvalidInput,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
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
