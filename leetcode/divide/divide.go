package divide

//Divide
//Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.
func Divide(dividend int, divisor int) int {
	divisorLtZero := false
	result := 0

	if divisor < 0 {
		divisorLtZero = true
		divisor = 0 - divisor
	}
	for dividend-divisor >= 0 {
		dividend = dividend - divisor
		result = result + 1
	}
	if !divisorLtZero {
		return result
	}
	return 0 - result
}
