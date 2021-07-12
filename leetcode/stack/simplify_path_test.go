package stack_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/leetcode/stack"
	"testing"
)

func TestSimplifyPath1(t *testing.T) {
	path := "/hello/"
	simplifiedPath := stack.SimplifyPath(path)
	expected := "/hello"

	if simplifiedPath != expected {
		t.Fatalf("Expected %v, received %v", expected, simplifiedPath)
	}
}

func TestSimplifyPath2(t *testing.T) {
	path := "/a/"
	simplifiedPath := stack.SimplifyPath(path)
	expected := "/a"

	if simplifiedPath != expected {
		t.Fatalf("Expected %v, received %v", expected, simplifiedPath)
	}
}

func TestSimplifyPath3(t *testing.T) {
	path := "/a/./b/../../c/"
	simplifiedPath := stack.SimplifyPath(path)
	expected := "/c"

	if simplifiedPath != expected {
		t.Fatalf("Expected %v, received %v", expected, simplifiedPath)
	}
}

func TestSimplifyPath4(t *testing.T) {
	path := "/a/../../b/../c/./"
	simplifiedPath := stack.SimplifyPath(path)
	expected := "/c"

	if simplifiedPath != expected {
		t.Fatalf("Expected %v, received %v", expected, simplifiedPath)
	}
}

func TestSimplifyPath5(t *testing.T) {
	path := "/a/b/c/d/././.."
	simplifiedPath := stack.SimplifyPath(path)
	expected := "/a/b/c"

	if simplifiedPath != expected {
		t.Fatalf("Expected %v, received %v", expected, simplifiedPath)
	}
}
