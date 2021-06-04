package stack

import (
	"errors"
	"github.com/SarthakMakhija/data-structures-and-algorithms/queue"
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

	if stack.stackTop == 0 {
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
			stack.Push(v)
		}
		if element == "+" || element == "*" || element == "-" || element == "/" {
			operand1 := stack.pop()
			operand2 := stack.pop()

			if operand1 == -1 || operand2 == -1 {
				return -1, errors.New("incorrect postfix expressions")
			}

			if element == "+" {
				stack.Push(operand1 + operand2)
			} else if element == "*" {
				stack.Push(operand1 * operand2)
			} else if element == "-" {
				stack.Push(operand2 - operand1)
			} else if element == "/" {
				stack.Push(operand2 / operand1)
			}
		}
	}
	return stack.stack[0], nil
}

//InFixToPostFix
//No parentheses, no minus
func InFixToPostFix(expression string) string {
	if len(expression) == 0 {
		return ""
	}

	operandStack := StringStack{}
	operatorStack := StringStack{}

	operatorPrecedence := func(operator string, other string) OperatorPrecedence {
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
			operandStack.push(element)
		}
		if element == "+" || element == "*" || element == "/" {
			topOperator := operatorStack.top()

			if len(topOperator) == 0 {
				operatorStack.push(element)
			} else {
				if operatorPrecedence(element, topOperator) == Eq || operatorPrecedence(element, topOperator) == Lt {
					operand2 := operandStack.pop()
					operand1 := operandStack.pop()
					operator := operatorStack.pop()

					operandStack.push(operand1 + operand2 + operator)
					operatorStack.push(element)
				}
				if operatorPrecedence(element, topOperator) == Gt {

					nextElement := string(expression[index+1])
					operatorStack.push(element)
					operandStack.push(nextElement)

					operand2 := operandStack.pop()
					operand1 := operandStack.pop()
					operator := operatorStack.pop()

					operandStack.push(operand1 + operand2 + operator)
					index = index + 1
				}
			}
		}
	}

	operand2 := operandStack.pop()
	operand1 := operandStack.pop()
	operator := operatorStack.pop()

	return operand1 + operand2 + operator
}

func (s *IntStack) CopyTo(destination *IntStack) {
	temporaryStack := IntStack{}
	for s.stackTop > 0 {
		temporaryStack.Push(s.pop())
	}

	for temporaryStack.stackTop > 0 {
		destination.Push(temporaryStack.pop())
	}
}

func (s *IntStack) All() []int {
	var allElements []int
	var top = s.stackTop - 1

	for top >= 0 {
		allElements = append(allElements, s.get(top))
		top = top - 1
	}
	return allElements
}

func DecimalToBinary(n int) string {
	stack := IntStack{}
	for n >= 1 {
		by2 := n >> 1
		remainder := n - by2*2
		n = by2
		stack.Push(remainder)
	}
	output := ""
	for stack.stackTop > 0 {
		output = output + strconv.Itoa(stack.pop())
	}
	return output
}

//Reverse
//Reverse a stack using a queue
func (s *IntStack) Reverse() {
	linearQueue := queue.LinearQueue{Size: len(s.stack)}
	for s.stackTop > 0 {
		pop := s.pop()
		_, _ = linearQueue.Add(pop)
	}
	allQueueElements := linearQueue.All()
	for index := 0; index < len(allQueueElements); index++ {
		s.Push(allQueueElements[index])
	}
}

type IntStack struct {
	stack    []int
	stackTop int
}

func (s *IntStack) Push(element int) {
	s.stack = append(s.stack, element)
	s.stackTop = s.stackTop + 1
}

func (s *IntStack) pop() int {
	s.stackTop = s.stackTop - 1
	if s.stackTop < 0 {
		return -1
	}
	top := s.stack[s.stackTop]
	s.stack = s.stack[0:s.stackTop]
	return top
}

func (s *IntStack) get(stackTop int) int {
	if stackTop < 0 {
		return -1
	}
	top := s.stack[stackTop]
	return top
}

type StringStack struct {
	stack    []string
	stackTop int
}

func (s *StringStack) push(element string) {
	s.stack = append(s.stack, element)
	s.stackTop = s.stackTop + 1
}

func (s *StringStack) pop() string {
	s.stackTop = s.stackTop - 1
	if s.stackTop < 0 {
		return ""
	}
	top := s.stack[s.stackTop]
	s.stack = s.stack[0:s.stackTop]
	return top
}

func (s *StringStack) top() string {
	stackTop := s.stackTop - 1
	if stackTop < 0 {
		return ""
	}
	return s.stack[stackTop]
}

type OperatorPrecedence int

const (
	Eq OperatorPrecedence = iota
	Gt
	Lt
	Unknown
)
