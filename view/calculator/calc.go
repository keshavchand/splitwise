package calculator

import (
	"fmt"
	"strconv"
	"strings"

	"errors"
)

func CalculatePostFix(input string) (float64, error) {
	stack := []float64{}
	args := strings.Split(input, " ")
	for _, arg := range args {
		switch arg {
		// TODO: Handle paranthesis
		case "+":
			num1, num2, nstack, err := GetTwo(stack)
			if err != nil {
				return 0, err
			}
			stack = append(nstack, num1+num2)
		case "-":
			num1, num2, nstack, err := GetTwo(stack)
			if err != nil {
				return 0, err
			}

			stack = append(nstack, num1-num2)
		case "/":
			num1, num2, nstack, err := GetTwo(stack)
			if err != nil {
				return 0, err
			}
			if num2 == 0 {
				return 0, errors.New("division by zero")
			}
			stack = append(nstack, num1/num2)
		case "*":
			num1, num2, nstack, err := GetTwo(stack)
			if err != nil {
				return 0, err
			}
			stack = append(nstack, num1*num2)
		default:
			number, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number %s", arg)
			}
			stack = append(stack, number)
		}
	}

	if len(stack) == 0 {
		return 0, errors.New("no numbers")
	}

	return stack[0], nil
}

func GetTwo(stack []float64) (float64, float64, []float64, error) {
	if len(stack) < 2 {
		return 0, 0, stack, errors.New("not enough numbers")
	}
	second := stack[len(stack)-1]
	first := stack[len(stack)-2]
	stack = stack[:len(stack)-2]
	return first, second, stack, nil
}
