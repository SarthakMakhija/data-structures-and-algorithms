package stringsops_test

import (
	stringsops "github.com/SarthakMakhija/data-structures-and-algorithms/strings_ops"
	"reflect"
	"testing"
)

func TestReplaceAndRemove1(t *testing.T) {
	characters := []rune{
		'a', 'c', 'd', 'b', 'b', 'c', 'a',
	}
	stringsops.ReplaceAndRemove(characters)
	expected := []rune{
		'd', 'd', 'c', 'd', 'c', 'd', 'd',
	}

	if !reflect.DeepEqual(characters, expected) {
		t.Fatalf("Expected %v, received %v", expected, characters)
	}
}

func TestReplaceAndRemove2(t *testing.T) {
	characters := []rune{
		'a', 'c', 'd', 'b',
	}
	stringsops.ReplaceAndRemove(characters)
	expected := []rune{
		'd', 'd', 'c', 'd',
	}

	if !reflect.DeepEqual(characters, expected) {
		t.Fatalf("Expected %v, received %v", expected, characters)
	}
}

func TestReplaceAndRemove3(t *testing.T) {
	var characters = make([]rune, 8)
	characters[0] = 'a'
	characters[1] = 'a'
	characters[2] = 'a'
	characters[3] = 'b'
	characters[4] = 'a'

	stringsops.ReplaceAndRemove(characters)
	expected := []rune{
		'd', 'd', 'd', 'd', 'd', 'd', 'd', 'd',
	}

	if !reflect.DeepEqual(characters, expected) {
		t.Fatalf("Expected %v, received %v", expected, characters)
	}
}
