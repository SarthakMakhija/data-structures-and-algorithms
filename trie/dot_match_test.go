package trie_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/trie"
	"reflect"
	"sort"
	"testing"
)

func TestAllWordsMatching1(t *testing.T) {
	node := trie.NewTrieNode()
	node.Add("hello")
	node.Add("help")
	node.Add("held")
	node.Add("how")
	node.Add("hot")

	allWords := node.AllWordsMatching("h.l.")
	sort.Strings(allWords)

	expected := []string{"help", "held"}
	sort.Strings(expected)

	if !reflect.DeepEqual(allWords, expected) {
		t.Fatalf("Expected %v, received %v", expected, allWords)
	}
}

func TestAllWordsMatching2(t *testing.T) {
	node := trie.NewTrieNode()
	node.Add("hello")
	node.Add("help")
	node.Add("held")
	node.Add("how")
	node.Add("hot")

	allWords := node.AllWordsMatching(".o.")
	sort.Strings(allWords)

	expected := []string{"how", "hot"}
	sort.Strings(expected)

	if !reflect.DeepEqual(allWords, expected) {
		t.Fatalf("Expected %v, received %v", expected, allWords)
	}
}

func TestAllWordsMatching3(t *testing.T) {
	node := trie.NewTrieNode()
	node.Add("hello")
	node.Add("help")
	node.Add("held")
	node.Add("how")
	node.Add("hot")

	allWords := node.AllWordsMatching("...")
	sort.Strings(allWords)

	expected := []string{"how", "hot"}
	sort.Strings(expected)

	if !reflect.DeepEqual(allWords, expected) {
		t.Fatalf("Expected %v, received %v", expected, allWords)
	}
}

func TestAllWordsMatching4(t *testing.T) {
	node := trie.NewTrieNode()
	node.Add("hello")
	node.Add("help")
	node.Add("held")
	node.Add("how")
	node.Add("hot")

	allWords := node.AllWordsMatching(".....")
	sort.Strings(allWords)

	expected := []string{"hello"}
	sort.Strings(expected)

	if !reflect.DeepEqual(allWords, expected) {
		t.Fatalf("Expected %v, received %v", expected, allWords)
	}
}

func TestAllWordsMatching5(t *testing.T) {
	node := trie.NewTrieNode()
	node.Add("attribute")
	node.Add("tribute")
	node.Add("tri")
	node.Add("but")

	allWords := node.AllWordsMatching(".....")
	sort.Strings(allWords)

	var expected []string

	if !reflect.DeepEqual(allWords, expected) {
		t.Fatalf("Expected %v, received %v", expected, allWords)
	}
}

func TestAllWordsMatching6(t *testing.T) {
	node := trie.NewTrieNode()
	node.Add("attribute")
	node.Add("tribute")
	node.Add("treat")
	node.Add("trout")

	allWords := node.AllWordsMatching("tr..t")
	sort.Strings(allWords)

	expected := []string{"treat", "trout"}
	sort.Strings(expected)

	if !reflect.DeepEqual(allWords, expected) {
		t.Fatalf("Expected %v, received %v", expected, allWords)
	}
}

func TestAllWordsMatching7(t *testing.T) {
	node := trie.NewTrieNode()
	node.Add("attribute")
	node.Add("tribute")
	node.Add("treat")
	node.Add("trout")

	allWords := node.AllWordsMatching("....t")
	sort.Strings(allWords)

	expected := []string{"treat", "trout"}
	sort.Strings(expected)

	if !reflect.DeepEqual(allWords, expected) {
		t.Fatalf("Expected %v, received %v", expected, allWords)
	}
}
