package stringsops_test

import (
	stringsops "github.com/SarthakMakhija/data-structures-and-algorithms/strings_ops"
	"testing"
)

func TestRunLengthEncode1(t *testing.T) {
	str := "aaaabcccaa"
	encoded := stringsops.RunLengthEncode(str)
	expected := "4a1b3c2a"

	if encoded != expected {
		t.Fatalf("Expected %v, received %v", expected, encoded)
	}
}

func TestRunLengthEncode2(t *testing.T) {
	str := "aaaabccca"
	encoded := stringsops.RunLengthEncode(str)
	expected := "4a1b3c1a"

	if encoded != expected {
		t.Fatalf("Expected %v, received %v", expected, encoded)
	}
}

func TestRunLengthEncode3(t *testing.T) {
	str := "abcd"
	encoded := stringsops.RunLengthEncode(str)
	expected := "1a1b1c1d"

	if encoded != expected {
		t.Fatalf("Expected %v, received %v", expected, encoded)
	}
}

func TestRunLengthDecode1(t *testing.T) {
	str := "4a1b3c2a"
	encoded := stringsops.RunLengthDecode(str)
	expected := "aaaabcccaa"

	if encoded != expected {
		t.Fatalf("Expected %v, received %v", expected, encoded)
	}
}

func TestRunLengthDecode2(t *testing.T) {
	str := "4a1b3c1a"
	encoded := stringsops.RunLengthDecode(str)
	expected := "aaaabccca"

	if encoded != expected {
		t.Fatalf("Expected %v, received %v", expected, encoded)
	}
}

func TestRunLengthDecode3(t *testing.T) {
	str := "1a1b1c1d"
	encoded := stringsops.RunLengthDecode(str)
	expected := "abcd"

	if encoded != expected {
		t.Fatalf("Expected %v, received %v", expected, encoded)
	}
}
