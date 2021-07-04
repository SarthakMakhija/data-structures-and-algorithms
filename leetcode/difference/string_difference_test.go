package difference_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/leetcode/difference"
	"testing"
)

func TestFindTheDifference1(t *testing.T) {
	result := difference.FindTheDifference("abcd", "abcde")
	differentCharacterAsString := string(result)

	if differentCharacterAsString != "e" {
		t.Fatalf("Expected %v, recevied %v", "e", differentCharacterAsString)
	}
}

func TestFindTheDifference2(t *testing.T) {
	result := difference.FindTheDifference("ae", "aea")
	differentCharacterAsString := string(result)

	if differentCharacterAsString != "a" {
		t.Fatalf("Expected %v, recevied %v", "a", differentCharacterAsString)
	}
}

func TestFindTheDifference3(t *testing.T) {
	result := difference.FindTheDifference("abcd", "bdcae")
	differentCharacterAsString := string(result)

	if differentCharacterAsString != "e" {
		t.Fatalf("Expected %v, recevied %v", "e", differentCharacterAsString)
	}
}

func TestFindTheDifference4(t *testing.T) {
	result := difference.FindTheDifference("", "y")
	differentCharacterAsString := string(result)

	if differentCharacterAsString != "y" {
		t.Fatalf("Expected %v, recevied %v", "y", differentCharacterAsString)
	}
}

func TestFindTheDifference5(t *testing.T) {
	result := difference.FindTheDifference("a", "aa")
	differentCharacterAsString := string(result)

	if differentCharacterAsString != "a" {
		t.Fatalf("Expected %v, recevied %v", "a", differentCharacterAsString)
	}
}

func TestFindTheDifference6(t *testing.T) {
	result := difference.FindTheDifference("aplep", "paelqp")
	differentCharacterAsString := string(result)

	if differentCharacterAsString != "q" {
		t.Fatalf("Expected %v, recevied %v", "q", differentCharacterAsString)
	}
}
