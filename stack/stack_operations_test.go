package stack_test

import (
	"reflect"
	"testing"

	"github.com/SarthakMakhija/data-structures-and-algorithms/stack"
)

func TestValidParenthesesMatching_1(t *testing.T) {

	doesParenthesesMatch := stack.MatchParentheses("(1+2)")
	if doesParenthesesMatch != true {
		t.Fatalf(`Expected true but found %v`, doesParenthesesMatch)
	}
}

func TestValidParenthesesMatching_2(t *testing.T) {

	doesParenthesesMatch := stack.MatchParentheses("(1+2+(3*3))")
	if doesParenthesesMatch != true {
		t.Fatalf(`Expected true but found %v`, doesParenthesesMatch)
	}
}

func TestParenthesesNotMatching_1(t *testing.T) {

	doesParenthesesMatch := stack.MatchParentheses("(1+2))")
	if doesParenthesesMatch != false {
		t.Fatalf(`Expected false but found %v`, doesParenthesesMatch)
	}
}

func TestParenthesesNotMatching_2(t *testing.T) {

	doesParenthesesMatch := stack.MatchParentheses("((1+2)")
	if doesParenthesesMatch != false {
		t.Fatalf(`Expected false but found %v`, doesParenthesesMatch)
	}
}

func TestParenthesesNotMatching_3(t *testing.T) {

	doesParenthesesMatch := stack.MatchParentheses("(1+2)+(3+4)+(9*8)))")
	if doesParenthesesMatch != false {
		t.Fatalf(`Expected false but found %v`, doesParenthesesMatch)
	}
}

func TestParenthesesNotMatching_4(t *testing.T) {

	doesParenthesesMatch := stack.MatchParentheses("((1+2")
	if doesParenthesesMatch != false {
		t.Fatalf(`Expected false but found %v`, doesParenthesesMatch)
	}
}

func TestEvaluatePostfixExpression_1(t *testing.T) {

	result, _ := stack.PostFixEvaluate("12+")
	if result != 3 {
		t.Fatalf(`Expected 3 but found %v`, result)
	}
}

func TestEvaluatePostfixExpression_2(t *testing.T) {

	result, _ := stack.PostFixEvaluate("12+3*")
	if result != 9 {
		t.Fatalf(`Expected 9 but found %v`, result)
	}
}

func TestEvaluatePostfixExpression_3(t *testing.T) {

	result, _ := stack.PostFixEvaluate("53+82-*")
	if result != 48 {
		t.Fatalf(`Expected 48 but found %v`, result)
	}
}

func TestEvaluatePostfixExpression_4(t *testing.T) {

	result, _ := stack.PostFixEvaluate("53+82-*6/")
	if result != 8 {
		t.Fatalf(`Expected 8 but found %v`, result)
	}
}

func TestEvaluatePostfixExpression_5(t *testing.T) {

	result, _ := stack.PostFixEvaluate("234*+5+")
	if result != 19 {
		t.Fatalf(`Expected 19 but found %v`, result)
	}
}

func TestEvaluateInvalidPostfixExpression(t *testing.T) {

	_, err := stack.PostFixEvaluate("+12")
	if err == nil {
		t.Fatal("Expected error but did not have any error")
	}
}

func TestInfixToPostfix_1(t *testing.T) {

	postfix := stack.InFixToPostFix("1+2")
	if postfix != "12+" {
		t.Fatalf("Expected 12+ but received %v", postfix)
	}
}

func TestInfixToPostfix_2(t *testing.T) {

	postfix := stack.InFixToPostFix("1+2+3")
	if postfix != "12+3+" {
		t.Fatalf("Expected 12+3+ but received %v", postfix)
	}
}

func TestInfixToPostfix_3(t *testing.T) {

	postfix := stack.InFixToPostFix("2+2*5")
	if postfix != "225*+" {
		t.Fatalf("Expected 225*+ but received %v", postfix)
	}
}

func TestInfixToPostfix_4(t *testing.T) {

	postfix := stack.InFixToPostFix("2+2*5/2")
	if postfix != "225*2/+" {
		t.Fatalf("Expected 225*2/+ but received %v", postfix)
	}
}

func TestInfixToPostfix_5(t *testing.T) {

	postfix := stack.InFixToPostFix("5*2*6")
	if postfix != "52*6*" {
		t.Fatalf("Expected 52*6* but received %v", postfix)
	}
}

func TestInfixToPostfix_6(t *testing.T) {

	postfix := stack.InFixToPostFix("5*2+6")
	if postfix != "52*6+" {
		t.Fatalf("Expected 52*6+ but received %v", postfix)
	}
}

func TestInfixToPostfix_7(t *testing.T) {

	postfix := stack.InFixToPostFix("5/2+6")
	if postfix != "52/6+" {
		t.Fatalf("Expected 52/6+ but received %v", postfix)
	}
}

func TestIntStackCopyTo_1(t *testing.T) {
	sourceStack := stack.IntStack{}
	sourceStack.Push(10)
	sourceStack.Push(20)
	sourceStack.Push(30)

	destinationStack := &stack.IntStack{}
	sourceStack.CopyTo(destinationStack)

	result := destinationStack.All()
	expected := []int{30, 20, 10}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

func TestIntStackCopyTo_2(t *testing.T) {
	sourceStack := stack.IntStack{}
	sourceStack.Push(10)
	sourceStack.Push(20)
	sourceStack.Push(30)
	sourceStack.Push(40)
	sourceStack.Push(50)

	destinationStack := &stack.IntStack{}
	sourceStack.CopyTo(destinationStack)

	result := destinationStack.All()
	expected := []int{50, 40, 30, 20, 10}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

func TestDecimalToBinary_1(t *testing.T) {
	output := stack.DecimalToBinary(10)
	if output != "1010" {
		t.Fatalf("Expected 1010 received %v", output)
	}
}

func TestDecimalToBinary_2(t *testing.T) {
	output := stack.DecimalToBinary(25)
	if output != "11001" {
		t.Fatalf("Expected 10001 received %v", output)
	}
}

func TestDecimalToBinary_3(t *testing.T) {
	output := stack.DecimalToBinary(9)
	if output != "1001" {
		t.Fatalf("Expected 1001 received %v", output)
	}
}

func TestDecimalToBinary_4(t *testing.T) {
	output := stack.DecimalToBinary(7)
	if output != "111" {
		t.Fatalf("Expected 111 received %v", output)
	}
}

func TestReverse_1(t *testing.T) {
	intStack := stack.IntStack{}
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	intStack.Push(40)
	intStack.Push(50)
	intStack.Push(60)

	intStack.Reverse()
	reversed := intStack.All()
	expected := []int{10, 20, 30, 40, 50, 60}

	if !reflect.DeepEqual(reversed, expected) {
		t.Fatalf("Expected %v received %v", expected, reversed)
	}
}

func TestReverse_2(t *testing.T) {
	intStack := stack.IntStack{}
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	intStack.Push(40)
	intStack.Push(50)
	intStack.Push(60)
	intStack.Push(80)

	intStack.Reverse()
	reversed := intStack.All()
	expected := []int{10, 20, 30, 40, 50, 60, 80}

	if !reflect.DeepEqual(reversed, expected) {
		t.Fatalf("Expected %v received %v", expected, reversed)
	}
}
