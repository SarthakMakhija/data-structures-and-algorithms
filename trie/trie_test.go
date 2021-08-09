package trie_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/trie"
	"reflect"
	"sort"
	"testing"
)

func TestAdd1(t *testing.T) {
	node := trie.NewTrieNode()
	node.Add("hello")
	node.Add("how")

	allWords := node.AllWords()
	expected := []string{"hello", "how"}

	if !reflect.DeepEqual(allWords, expected) {
		t.Fatalf("Expected all words to be %v, received %v", expected, allWords)
	}
}

func TestAdd2(t *testing.T) {
	node := trie.NewTrieNode()
	node.Add("camera")
	node.Add("came")
	node.Add("card")

	allWords := node.AllWords()
	expected := []string{"came", "camera", "card"}

	if !reflect.DeepEqual(allWords, expected) {
		t.Fatalf("Expected all words to be %v, received %v", expected, allWords)
	}
}

func TestAdd3(t *testing.T) {
	node := trie.NewTrieNode()
	node.Add("abc")
	node.Add("abgl")
	node.Add("abcd")
	node.Add("lmo")
	node.Add("lmno")

	allWords := node.AllWords()
	sort.Strings(allWords)

	expected := []string{"abc", "abcd", "abgl", "lmo", "lmno"}
	sort.Strings(expected)

	if !reflect.DeepEqual(allWords, expected) {
		t.Fatalf("Expected all words to be %v, received %v", expected, allWords)
	}
}
