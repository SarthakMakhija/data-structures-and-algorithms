package stack_test

import (
	"testing"

	"example.com/data-structures-and-algorithms/stack"
)

func TestValidParenthesesMatching_1(t *testing.T) {

	does_parentheses_match := stack.MatchParentheses("(1+2)")
	if does_parentheses_match != true {
		t.Fatalf(`Expected true but found %v`, does_parentheses_match)
	}
}

func TestValidParenthesesMatching_2(t *testing.T) {

	does_parentheses_match := stack.MatchParentheses("(1+2+(3*3))")
	if does_parentheses_match != true {
		t.Fatalf(`Expected true but found %v`, does_parentheses_match)
	}
}

func TestParenthesesNotMatching_1(t *testing.T) {

	does_parentheses_match := stack.MatchParentheses("(1+2))")
	if does_parentheses_match != false {
		t.Fatalf(`Expected false but found %v`, does_parentheses_match)
	}
}

func TestParenthesesNotMatching_2(t *testing.T) {

	does_parentheses_match := stack.MatchParentheses("((1+2)")
	if does_parentheses_match != false {
		t.Fatalf(`Expected false but found %v`, does_parentheses_match)
	}
}

func TestParenthesesNotMatching_3(t *testing.T) {

	does_parentheses_match := stack.MatchParentheses("(1+2)+(3+4)+(9*8)))")
	if does_parentheses_match != false {
		t.Fatalf(`Expected false but found %v`, does_parentheses_match)
	}
}

func TestParenthesesNotMatching_4(t *testing.T) {

	does_parentheses_match := stack.MatchParentheses("((1+2")
	if does_parentheses_match != false {
		t.Fatalf(`Expected false but found %v`, does_parentheses_match)
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
