package stringsops_test

import (
	stringsops "github.com/SarthakMakhija/data-structures-and-algorithms/strings_ops"
	"testing"
)

func TestIsPalindromic1(t *testing.T) {
	str := "A man, a plan, a canal, Panama."
	result := stringsops.IsPalindromic(str)

	if result != true {
		t.Fatalf("%v should be palindromic", str)
	}
}

func TestIsPalindromic2(t *testing.T) {
	str := "Able was I, ere I saw Elba!"
	result := stringsops.IsPalindromic(str)

	if result != true {
		t.Fatalf("%v should be palindromic", str)
	}
}

func TestIsPalindromic3(t *testing.T) {
	str := "Ray a Ray"
	result := stringsops.IsPalindromic(str)

	if result != false {
		t.Fatalf("%v should not be palindromic", str)
	}
}
