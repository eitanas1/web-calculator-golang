package calc

import (
	"errors"
	"strconv"
)

var ErrorInvalidInput = errors.New("expression is not valid")
var ErrorDevisionByZero = errors.New("division by zero is not allowed")

func Calc(expression string) (float64, error) {
	if expression == "" {
		return 0, ErrorInvalidInput
	}
	priority := map[string]int{
		"(": 0,
		")": 1,
		"+": 2,
		"-": 2,
		"*": 3,
		"/": 3,
	}
	stack := []string{}
	stackResultPolish := make([]float64, 0)
	reversePolishNotation := ""

	for _, symbol := range expression {
		if _, ok := priority[string(symbol)]; ok {
			if string(symbol) == ")" {
				for i := len(stack) - 1; i >= 0 && stack[i] != "("; i-- {
					reversePolishNotation += string(stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				if stack[len(stack)-1] == "(" {
					stack = stack[:len(stack)-1]
				}
				continue
			}
			for len(stack) > 0 && priority[stack[len(stack)-1]] >= priority[string(symbol)] && string(symbol) != "(" {
				reversePolishNotation += stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, string(symbol))
		} else if symbol >= '0' && symbol <= '9' {
			reversePolishNotation += string(symbol)
		} else {
			return 0, ErrorInvalidInput
		}
	}
	for len(stack) > 0 {
		reversePolishNotation += stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}

	for _, symbol := range reversePolishNotation {
		floatSymbol, err := strconv.ParseFloat(string(symbol), 64)
		if err == nil {
			stackResultPolish = append(stackResultPolish, floatSymbol)
		} else {
			if len(stackResultPolish) < 2 {
				return 0, ErrorInvalidInput
			}
			num1 := stackResultPolish[len(stackResultPolish)-1]
			num2 := stackResultPolish[len(stackResultPolish)-2]
			stackResultPolish = stackResultPolish[:len(stackResultPolish)-2]

			switch string(symbol) {
			case "+":
				stackResultPolish = append(stackResultPolish, num1+num2)
			case "-":
				stackResultPolish = append(stackResultPolish, num2-num1)
			case "*":
				stackResultPolish = append(stackResultPolish, num2*num1)
			case "/":
				if num1 == 0 {
					return 0, ErrorDevisionByZero
				}
				stackResultPolish = append(stackResultPolish, num2/num1)
			}
		}
	}
	return stackResultPolish[0], nil
}
