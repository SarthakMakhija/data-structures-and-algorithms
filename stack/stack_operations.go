package stack

import (
	"errors"
	"strconv"
)

func MatchParentheses(expression string) bool {
	if len(expression) == 0 {
		return false
	}

	stack := StringStack{}
	for index := 0; index < len(expression); index++ {
		element := string(expression[index])
		if element == "(" {
			stack.push(element)
		} else if element == ")" {
			stack.pop()
		}
	}

	if stack.stack_top == 0 {
		return true
	} else {
		return false
	}
}

func PostFixEvaluate(expression string) (int, error) {
	if len(expression) == 0 {
		return 0, nil
	}

	stack := IntStack{}
	for index := 0; index < len(expression); index++ {
		element := string(expression[index])
		if v, err := strconv.Atoi(element); err == nil {
			stack.push(v)
		}
		if element == "+" || element == "*" || element == "-" || element == "/" {
			operand1 := stack.pop()
			operand2 := stack.pop()

			if operand1 == -1 || operand2 == -1 {
				return -1, errors.New("incorrect postfix expressions")
			}

			if element == "+" {
				stack.push(operand1 + operand2)
			} else if element == "*" {
				stack.push(operand1 * operand2)
			} else if element == "-" {
				stack.push(operand2 - operand1)
			} else if element == "/" {
				stack.push(operand2 / operand1)
			}
		}
	}
	return stack.stack[0], nil
}

type IntStack struct {
	stack     []int
	stack_top int
}

func (s *IntStack) push(element int) {
	s.stack = append(s.stack, element)
	s.stack_top = s.stack_top + 1
}

func (s *IntStack) pop() int {
	s.stack_top = s.stack_top - 1
	if s.stack_top < 0 {
		return -1
	}
	top := s.stack[s.stack_top]
	s.stack = s.stack[0:s.stack_top]
	return top
}

type StringStack struct {
	stack     []string
	stack_top int
}

func (s *StringStack) push(element string) {
	s.stack = append(s.stack, element)
	s.stack_top = s.stack_top + 1
}

func (s *StringStack) pop() string {
	s.stack_top = s.stack_top - 1
	if s.stack_top < 0 {
		return ""
	}
	top := s.stack[s.stack_top]
	s.stack = s.stack[0:s.stack_top]
	return top
}
