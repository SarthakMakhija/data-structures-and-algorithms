package stack

func MatchParentheses(expression string) bool {
	if len(expression) == 0 {
		return false
	}

	var stack_push func(rune)
	var stack_pop func()

	var stack_top = 0
	var stack []rune

	stack_push = func(element rune) {
		stack = append(stack, element)
		stack_top = stack_top + 1
	}

	stack_pop = func() {
		stack_top = stack_top - 1
		if stack_top < 0 {
			return
		}
		stack = stack[0:stack_top]
	}

	for index := 0; index < len(expression); index++ {
		char := rune(expression[index])
		if char == '(' {
			stack_push(char)
		} else if char == ')' {
			stack_pop()
		}
	}

	if stack_top == 0 {
		return true
	} else {
		return false
	}
}
