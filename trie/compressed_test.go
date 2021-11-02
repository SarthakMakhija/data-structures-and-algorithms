package trie_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/trie"
	"testing"
)

func TestFindsAPatternInAStringUsingCompressedTrie1(t *testing.T) {
	compressedTrie := trie.NewCompressedTrieNaive("banana")
	contains := compressedTrie.Contains("nan")

	if contains != true {
		t.Fatalf("Expected nan to be contained but was not")
	}
}

func TestFindsAPatternInAStringUsingCompressedTrie2(t *testing.T) {
	compressedTrie := trie.NewCompressedTrieNaive("banana")
	contains := compressedTrie.Contains("ban")

	if contains != true {
		t.Fatalf("Expected ban to be contained but was not")
	}
}

func TestFindsAPatternInAStringUsingCompressedTrie3(t *testing.T) {
	compressedTrie := trie.NewCompressedTrieNaive("banana")
	contains := compressedTrie.Contains("nana")

	if contains != true {
		t.Fatalf("Expected nana to be contained but was not")
	}
}

func TestFindsAPatternInAStringUsingCompressedTrie4(t *testing.T) {
	compressedTrie := trie.NewCompressedTrieNaive("banana")
	contains := compressedTrie.Contains("banana")

	if contains != true {
		t.Fatalf("Expected banana to be contained but was not")
	}
}

func TestFindsAPatternInAStringUsingCompressedTrie5(t *testing.T) {
	compressedTrie := trie.NewCompressedTrieNaive("banana")
	contains := compressedTrie.Contains("a")

	if contains != true {
		t.Fatalf("Expected a to be contained but was not")
	}
}

func TestDoesNotFindAPatternInAStringUsingCompressedTrie6(t *testing.T) {
	compressedTrie := trie.NewCompressedTrieNaive("banana")
	contains := compressedTrie.Contains("anc")

	if contains != false {
		t.Fatalf("Expected anc to not be contained but was")
	}
}

func TestFindsASuffixInAStringUsingCompressedTrie1(t *testing.T) {
	compressedTrie := trie.NewCompressedTrieNaive("banana")
	contains := compressedTrie.ContainsSuffix("ana")

	if contains != true {
		t.Fatalf("Expected ana suffix to be contained but was not")
	}
}

func TestFindsASuffixInAStringUsingCompressedTrie2(t *testing.T) {
	compressedTrie := trie.NewCompressedTrieNaive("banana")
	contains := compressedTrie.ContainsSuffix("nana")

	if contains != true {
		t.Fatalf("Expected nana suffix to be contained but was not")
	}
}

func TestFindsASuffixInAStringUsingCompressedTrie3(t *testing.T) {
	compressedTrie := trie.NewCompressedTrieNaive("banana")
	contains := compressedTrie.ContainsSuffix("banana")

	if contains != true {
		t.Fatalf("Expected banana suffix to be contained but was not")
	}
}

func TestDoesNotFindASuffixInAStringUsingCompressedTrie(t *testing.T) {
	compressedTrie := trie.NewCompressedTrieNaive("banana")
	contains := compressedTrie.ContainsSuffix("nan")

	if contains != false {
		t.Fatalf("Expected nan suffix to not be contained but was")
	}
}
