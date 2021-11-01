package trie_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/trie"
	"testing"
)

func TestFindsAPatternInAStringUsingSuffixTrie1(t *testing.T) {
	suffixTrie := trie.NewSuffixTrie("banana")
	contains := suffixTrie.Contains("nan")

	if contains != true {
		t.Fatalf("Expected nan to be contained but was not")
	}
}

func TestFindsAPatternInAStringUsingSuffixTrie2(t *testing.T) {
	suffixTrie := trie.NewSuffixTrie("banana")
	contains := suffixTrie.Contains("ban")

	if contains != true {
		t.Fatalf("Expected ban to be contained but was not")
	}
}

func TestFindsAPatternInAStringUsingSuffixTrie3(t *testing.T) {
	suffixTrie := trie.NewSuffixTrie("banana")
	contains := suffixTrie.Contains("nana")

	if contains != true {
		t.Fatalf("Expected nana to be contained but was not")
	}
}

func TestFindsAPatternInAStringUsingSuffixTrie4(t *testing.T) {
	suffixTrie := trie.NewSuffixTrie("banana")
	contains := suffixTrie.Contains("banana")

	if contains != true {
		t.Fatalf("Expected banana to be contained but was not")
	}
}

func TestFindsAPatternInAStringUsingSuffixTrie5(t *testing.T) {
	suffixTrie := trie.NewSuffixTrie("banana")
	contains := suffixTrie.Contains("a")

	if contains != true {
		t.Fatalf("Expected a to be contained but was not")
	}
}

func TestDoesNotFindAPatternInAStringUsingSuffixTrie1(t *testing.T) {
	suffixTrie := trie.NewSuffixTrie("banana")
	contains := suffixTrie.Contains("anc")

	if contains != false {
		t.Fatalf("Expected anc to not be contained but was")
	}
}

func TestFindsASuffixInAStringUsingSuffixTrie1(t *testing.T) {
	suffixTrie := trie.NewSuffixTrie("banana")
	contains := suffixTrie.ContainsSuffix("ana")

	if contains != true {
		t.Fatalf("Expected ana suffix to be contained but was not")
	}
}

func TestFindsASuffixInAStringUsingSuffixTrie2(t *testing.T) {
	suffixTrie := trie.NewSuffixTrie("banana")
	contains := suffixTrie.ContainsSuffix("nana")

	if contains != true {
		t.Fatalf("Expected nana suffix to be contained but was not")
	}
}

func TestFindsASuffixInAStringUsingSuffixTrie3(t *testing.T) {
	suffixTrie := trie.NewSuffixTrie("banana")
	contains := suffixTrie.ContainsSuffix("banana")

	if contains != true {
		t.Fatalf("Expected banana suffix to be contained but was not")
	}
}

func TestDoesNotFindASuffixInAStringUsingSuffixTrie1(t *testing.T) {
	suffixTrie := trie.NewSuffixTrie("banana")
	contains := suffixTrie.ContainsSuffix("nan")

	if contains != false {
		t.Fatalf("Expected nan suffix to not be contained but was")
	}
}
