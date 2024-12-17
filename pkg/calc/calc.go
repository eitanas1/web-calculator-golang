package calc

import (
	"strconv"
)

func Calc(stringExpression string) (float64, error) {
	expression, err := tokenize(stringExpression)
	if err != nil {
		return 0, err
	}

	if len(expression) == 0 {
		return 0, ErrorEmptyExpression
	}

	priority := map[string]int{
		"(": 0,
		")": 1,
		"+": 2,
		"-": 2,
		"*": 3,
		"/": 3,
	}
	stack := []Token{}
	stackResultPolish := make([]float64, 0)
	reversePolishNotation := []Token{}

	for _, token := range expression {
		if _, ok := priority[token.Value]; ok {
			if token.Value == ")" {
				for i := len(stack) - 1; i >= 0 && stack[i].Value != "("; i-- {
					reversePolishNotation = append(reversePolishNotation, lastToken(stack))
					stack = stack[:len(stack)-1]
				}

				if len(stack) > 0 && stack[len(stack)-1].Value == "(" {
					stack = stack[:len(stack)-1]
				}
				continue
			}

			for len(stack) > 0 && priority[stack[len(stack)-1].Value] >= priority[token.Value] && token.Value != "(" {
				reversePolishNotation = append(reversePolishNotation, lastToken(stack))
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)

		} else if token.IsNumber {
			reversePolishNotation = append(reversePolishNotation, lastToken(stack))
		} else {
			return 0, ErrorInvalidInput
		}
	}

	for len(stack) > 0 {
		reversePolishNotation = append(reversePolishNotation, lastToken(stack))
		stack = stack[:len(stack)-1]
	}

	for _, token := range reversePolishNotation {
		floatNumber, err := strconv.ParseFloat(token.Value, 64)

		if err == nil {
			stackResultPolish = append(stackResultPolish, floatNumber)
		} else {
			if len(stackResultPolish) < 2 {
				return 0, ErrorInvalidInput
			}

			num1 := stackResultPolish[len(stackResultPolish)-1]
			num2 := stackResultPolish[len(stackResultPolish)-2]
			stackResultPolish = stackResultPolish[:len(stackResultPolish)-2]

			switch token.Value {
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
