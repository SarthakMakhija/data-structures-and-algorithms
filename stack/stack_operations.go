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

//No parentheses, no minus
func InFixToPostFix(expression string) string {
	if len(expression) == 0 {
		return ""
	}

	operand_stack := StringStack{}
	operator_stack := StringStack{}

	operator_precedence := func(operator string, other string) OperatorPrecedence {
		if operator == other {
			return Eq
		}
		if operator == "*" && other == "+" {
			return Gt
		}
		if operator == "/" && other == "+" {
			return Gt
		}
		if operator == "+" && other == "*" {
			return Lt
		}
		if operator == "+" && other == "/" {
			return Lt
		}
		return Unknown
	}

	for index := 0; index < len(expression); index++ {
		element := string(expression[index])
		if _, err := strconv.Atoi(element); err == nil {
			operand_stack.push(element)
		}
		if element == "+" || element == "*" || element == "/" {
			top_operator := operator_stack.top()

			if len(top_operator) == 0 {
				operator_stack.push(element)
			} else {
				if operator_precedence(element, top_operator) == Eq || operator_precedence(element, top_operator) == Lt {
					operand2 := operand_stack.pop()
					operand1 := operand_stack.pop()
					operator := operator_stack.pop()

					operand_stack.push(operand1 + operand2 + operator)
					operator_stack.push(element)
				}
				if operator_precedence(element, top_operator) == Gt {

					next_element := string(expression[index+1])
					operator_stack.push(element)
					operand_stack.push(next_element)

					operand2 := operand_stack.pop()
					operand1 := operand_stack.pop()
					operator := operator_stack.pop()

					operand_stack.push(operand1 + operand2 + operator)
					index = index + 1
				}
			}
		}
	}

	operand2 := operand_stack.pop()
	operand1 := operand_stack.pop()
	operator := operator_stack.pop()

	return operand1 + operand2 + operator
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

func (s *StringStack) top() string {
	stack_top := s.stack_top - 1
	if stack_top < 0 {
		return ""
	}
	return s.stack[stack_top]
}

type OperatorPrecedence int

const (
	Eq OperatorPrecedence = iota
	Gt
	Lt
	Unknown
)
