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
