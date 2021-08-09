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
	sort.Strings(allWords)

	expected := []string{"hello", "how"}
	sort.Strings(expected)

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
	sort.Strings(allWords)

	expected := []string{"came", "camera", "card"}
	sort.Strings(expected)

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

func TestExistsCompleteWord1(t *testing.T) {
	node := trie.NewTrieNode()
	node.Add("hello")
	node.Add("hail")

	exists := node.ExistsCompleteWord("hello")

	if exists != true {
		t.Fatalf("Expected hello to exist but found %v", exists)
	}
}

func TestExistsCompleteWord2(t *testing.T) {
	node := trie.NewTrieNode()
	node.Add("hellos")

	exists := node.ExistsCompleteWord("hello")

	if exists != false {
		t.Fatalf("Expected hello to not exist but found %v", exists)
	}
}

func TestExistsCompleteWord3(t *testing.T) {
	node := trie.NewTrieNode()
	node.Add("camera")
	node.Add("came")
	node.Add("card")

	exists := node.ExistsCompleteWord("card")

	if exists != true {
		t.Fatalf("Expected card to exist but found %v", exists)
	}
}

func TestExistsCompleteWord4(t *testing.T) {
	node := trie.NewTrieNode()
	node.Add("camera")
	node.Add("came")
	node.Add("card")

	exists := node.ExistsCompleteWord("cards")

	if exists != false {
		t.Fatalf("Expected cards to not exist but found %v", exists)
	}
}

func TestExistsPrefix1(t *testing.T) {
	node := trie.NewTrieNode()
	node.Add("camera")
	node.Add("came")

	exists := node.ExistsPrefix("cam")

	if exists != true {
		t.Fatalf("Expected prefix cam to exist but found %v", exists)
	}
}

func TestExistsPrefix2(t *testing.T) {
	node := trie.NewTrieNode()
	node.Add("hellos")
	node.Add("hail")

	exists := node.ExistsPrefix("hell")

	if exists != true {
		t.Fatalf("Expected prefix hell to exist but found %v", exists)
	}
}

func TestExistsPrefix3(t *testing.T) {
	node := trie.NewTrieNode()
	node.Add("hellos")
	node.Add("hail")

	exists := node.ExistsPrefix("heli")

	if exists != false {
		t.Fatalf("Expected prefix heli to exist but found %v", exists)
	}
}

func TestAutoCompleteWithPrefix1(t *testing.T) {
	node := trie.NewTrieNode()
	node.Add("hellos")
	node.Add("hello")
	node.Add("helicopter")
	node.Add("help")
	node.Add("held")

	words := node.AutoCompleteWithPrefix("hel")
	sort.Strings(words)

	expected := []string{"hellos", "hello", "helicopter", "help", "held"}
	sort.Strings(expected)

	if !reflect.DeepEqual(words, expected) {
		t.Fatalf("Expected autocompleted words to be %v, received %v", expected, words)
	}
}

func TestAutoCompleteWithPrefix2(t *testing.T) {
	node := trie.NewTrieNode()
	node.Add("technical")
	node.Add("techno")
	node.Add("tech debt")
	node.Add("tech fest")

	words := node.AutoCompleteWithPrefix("tech")
	sort.Strings(words)

	expected := []string{"technical", "techno", "tech debt", "tech fest"}
	sort.Strings(expected)

	if !reflect.DeepEqual(words, expected) {
		t.Fatalf("Expected autocompleted words to be %v, received %v", expected, words)
	}
}
