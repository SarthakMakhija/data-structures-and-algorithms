package stack_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/stack"
	"testing"
)

func TestEvaluatePrefix1(t *testing.T) {
	result := stack.EvaluatePrefix("+25")
	if result != 7 {
		t.Fatalf("Expected 7, received %v", result)
	}
}

func TestEvaluatePrefix2(t *testing.T) {
	result := stack.EvaluatePrefix("*93")
	if result != 27 {
		t.Fatalf("Expected 27, received %v", result)
	}
}

func TestEvaluatePrefix3(t *testing.T) {
	result := stack.EvaluatePrefix("-93")
	if result != 6 {
		t.Fatalf("Expected 6, received %v", result)
	}
}

func TestEvaluatePrefix4(t *testing.T) {
	result := stack.EvaluatePrefix("/93")
	if result != 3 {
		t.Fatalf("Expected 3, received %v", result)
	}
}

func TestEvaluatePrefix5(t *testing.T) {
	result := stack.EvaluatePrefix("*+543")
	if result != 27 {
		t.Fatalf("Expected 27, received %v", result)
	}
}

func TestEvaluatePrefix6(t *testing.T) {
	result := stack.EvaluatePrefix("+*72*52")
	if result != 24 {
		t.Fatalf("Expected 24, received %v", result)
	}
}

func TestEvaluatePrefix7(t *testing.T) {
	result := stack.EvaluatePrefix("+8*65")
	if result != 38 {
		t.Fatalf("Expected 38, received %v", result)
	}
}
