package copy_test

import (
	"github.com/SarthakMakhija/data-structures-and-algorithms/leetcode/copy"
	"reflect"
	"testing"
)

func TestCopyWithRandomPointer1(t *testing.T) {
	seven := copy.Node{
		Val:    7,
		Random: nil,
	}
	one := copy.Node{
		Val:    1,
		Random: &seven,
	}
	eleven := copy.Node{
		Val: 11,
	}
	ten := copy.Node{
		Val:    10,
		Next:   &one,
		Random: &eleven,
	}
	thirteen := copy.Node{
		Val:    13,
		Next:   &eleven,
		Random: &seven,
	}
	seven.Next = &thirteen
	eleven.Next = &ten
	eleven.Random = &one

	copied := copy.CopyWithRandomPointer(&seven)
	output := copied.All()
	expected := []copy.NodeValue{
		{
			Value:       7,
			RandomValue: 0,
		},
		{
			Value:       13,
			RandomValue: 7,
		},
		{
			Value:       11,
			RandomValue: 1,
		},
		{
			Value:       10,
			RandomValue: 11,
		},
		{
			Value:       1,
			RandomValue: 7,
		},
	}
	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}

func TestCopyWithRandomPointer2(t *testing.T) {
	two := copy.Node{
		Val: 2,
	}
	one := copy.Node{
		Val:    1,
		Next:   &two,
		Random: &two,
	}
	two.Random = &two

	copied := copy.CopyWithRandomPointer(&one)
	output := copied.All()
	expected := []copy.NodeValue{
		{
			Value:       1,
			RandomValue: 2,
		},
		{
			Value:       2,
			RandomValue: 2,
		},
	}
	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}

func TestCopyWithRandomPointer3(t *testing.T) {
	thirdThree := copy.Node{
		Val: 3,
	}
	secondThree := copy.Node{
		Val:  3,
		Next: &thirdThree,
	}
	firstThree := copy.Node{
		Val:  3,
		Next: &secondThree,
	}
	secondThree.Random = &firstThree

	copied := copy.CopyWithRandomPointer(&firstThree)
	output := copied.All()
	expected := []copy.NodeValue{
		{
			Value:       3,
			RandomValue: 0,
		},
		{
			Value:       3,
			RandomValue: 3,
		},
		{
			Value:       3,
			RandomValue: 0,
		},
	}
	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}
