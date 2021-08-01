package stringsops_test

import (
	stringsops "github.com/SarthakMakhija/data-structures-and-algorithms/strings_ops"
	"testing"
)

func TestLookAndSayValueForK2(t *testing.T) {
	k := 2
	result := stringsops.LookAndSayValueFor(k)
	expected := "11"

	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestLookAndSayValueForK3(t *testing.T) {
	k := 3
	result := stringsops.LookAndSayValueFor(k)
	expected := "21"

	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestLookAndSayValueForK4(t *testing.T) {
	k := 4
	result := stringsops.LookAndSayValueFor(k)
	expected := "1211"

	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestLookAndSayValueForK5(t *testing.T) {
	k := 5
	result := stringsops.LookAndSayValueFor(k)
	expected := "111221"

	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestLookAndSayValueForK6(t *testing.T) {
	k := 6
	result := stringsops.LookAndSayValueFor(k)
	expected := "Im312211"

	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}
