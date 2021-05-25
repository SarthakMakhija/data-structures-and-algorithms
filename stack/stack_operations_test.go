package stack_test

import (
	"testing"

	"example.com/data-structures-and-algorithms/stack"
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
