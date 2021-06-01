package linkedlist_test

import (
	linkedlist "github.com/SarthakMakhija/data-structures-and-algorithms/linked_list"
	"strings"
	"testing"
)

func TestPolynomialListDisplay_1(t *testing.T) {
	list := linkedlist.PolynomialList{}
	list.Add(10, 4)
	list.Add(10, 1)
	list.Add(5, 0)
	list.Add(20, 2)
	list.Add(30, 3)

	expression := strings.Trim(list.Expression(), " ")
	expected := "10x4 + 30x3 + 20x2 + 10x1 + 5x0"

	if expression != expected {
		t.Fatalf("Expecting %v, received %v", expected, expression)
	}
}

func TestPolynomialListDisplay_2(t *testing.T) {
	list := linkedlist.PolynomialList{}
	list.Add(5, 0)
	list.Add(20, 2)
	list.Add(30, 3)
	list.Add(50, 2)
	list.Add(50, 3)

	expression := strings.Trim(list.Expression(), " ")
	expected := "80x3 + 70x2 + 5x0"

	if expression != expected {
		t.Fatalf("Expecting %v, received %v", expected, expression)
	}
}

func TestPolynomialListPlus_1(t *testing.T) {
	firstList := linkedlist.PolynomialList{}
	firstList.Add(10, 1)
	firstList.Add(5, 0)
	firstList.Add(20, 2)
	firstList.Add(30, 3)

	secondList := linkedlist.PolynomialList{}
	secondList.Add(10, 1)
	secondList.Add(5, 0)
	secondList.Add(2, 2)
	secondList.Add(50, 3)

	resultList := firstList.Plus(secondList)

	expression := strings.Trim(resultList.Expression(), " ")
	expected := "80x3 + 22x2 + 20x1 + 10x0"

	if expression != expected {
		t.Fatalf("Expecting %v, received %v", expected, expression)
	}
}

func TestPolynomialListPlus_2(t *testing.T) {
	firstList := linkedlist.PolynomialList{}
	firstList.Add(10, 1)
	firstList.Add(5, 2)
	firstList.Add(5, 3)
	firstList.Add(20, 2)
	firstList.Add(30, 3)

	secondList := linkedlist.PolynomialList{}
	secondList.Add(10, 1)
	secondList.Add(5, 0)
	secondList.Add(50, 3)

	resultList := firstList.Plus(secondList)

	expression := strings.Trim(resultList.Expression(), " ")
	expected := "85x3 + 25x2 + 20x1 + 5x0"

	if expression != expected {
		t.Fatalf("Expecting %v, received %v", expected, expression)
	}
}
