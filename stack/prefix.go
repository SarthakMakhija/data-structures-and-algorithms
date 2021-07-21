package stack

func EvaluatePrefix(expression string) int {
	if len(expression) == 0 {
		return -1
	}
	const unicodeZero = 48
	operatorStack := StringStack{}
	operandStack := IntStack{}
	operandCount := 0

	resetOperandCount := func() {
		operandCount = 0
	}
	evaluate := func() {
		operand1 := operandStack.Pop()
		operand2 := operandStack.Pop()
		switch operatorStack.Pop() {
		case "+":
			operandStack.Push(operand2 + operand1)
		case "-":
			operandStack.Push(operand2 - operand1)
		case "*":
			operandStack.Push(operand2 * operand1)
		case "/":
			operandStack.Push(operand2 / operand1)
		}
	}
	areTwoOperandsAvailable := func() bool {
		if operandCount == 2 {
			return true
		}
		return false
	}
	for _, ch := range expression {
		switch string(ch) {
		case "+", "-", "/", "*":
			operatorStack.Push(string(ch))
			resetOperandCount()
		default:
			operandStack.Push(int(ch) - unicodeZero)
			operandCount = operandCount + 1
			if areTwoOperandsAvailable() {
				resetOperandCount()
				evaluate()
			}
		}
	}

	for !operatorStack.IsEmpty() {
		evaluate()
	}
	return operandStack.Pop()
}
