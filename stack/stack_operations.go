package stack

import (
	"errors"
	"strconv"
)

func MatchParentheses(expression string) bool {
	if len(expression) == 0 {
		return false
	}

	var stack_push func(string)
	var stack_pop func() string

	var stack_top = 0
	var stack []string

	stack_push = func(element string) {
		stack = append(stack, element)
		stack_top = stack_top + 1
	}

	stack_pop = func() string {
		stack_top = stack_top - 1
		if stack_top < 0 {
			return ""
		}
		top := stack[stack_top]
		stack = stack[0:stack_top]
		return top
	}

	for index := 0; index < len(expression); index++ {
		element := string(expression[index])
		if element == "(" {
			stack_push(element)
		} else if element == ")" {
			stack_pop()
		}
	}

	if stack_top == 0 {
		return true
	} else {
		return false
	}
}

func PostFixEvaluate(expression string) (int, error) {
	if len(expression) == 0 {
		return 0, nil
	}

	var stack_push func(int)
	var stack_pop func() int

	var stack_top = 0
	var stack []int

	stack_push = func(element int) {
		stack = append(stack, element)
		stack_top = stack_top + 1
	}

	stack_pop = func() int {
		stack_top = stack_top - 1
		if stack_top < 0 {
			return -1
		}
		top := stack[stack_top]
		stack = stack[0:stack_top]
		return top
	}

	for index := 0; index < len(expression); index++ {
		element := string(expression[index])
		if v, err := strconv.Atoi(element); err == nil {
			stack_push(v)
		}
		if element == "+" || element == "*" || element == "-" || element == "/" {
			operand1 := stack_pop()
			operand2 := stack_pop()
			if operand1 == -1 || operand2 == -1 {
				return -1, errors.New("incorrect postfix expressions")
			}

			if element == "+" {
				stack_push(operand1 + operand2)
			} else if element == "*" {
				stack_push(operand1 * operand2)
			} else if element == "-" {
				stack_push(operand2 - operand1)
			} else if element == "/" {
				stack_push(operand2 / operand1)
			}
		}
	}
	return stack[0], nil
}
