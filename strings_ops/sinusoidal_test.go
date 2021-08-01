package stringsops_test

import (
	stringsops "github.com/SarthakMakhija/data-structures-and-algorithms/strings_ops"
	"testing"
)

func TestSinusoidalFrom1(t *testing.T) {

	str := "Hello_World!"
	output := stringsops.SinusoidalFrom(str)
	expected := "e_lHloWrdlo!"

	if output != expected {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}

func TestSinusoidalFrom2(t *testing.T) {

	str := "Hello"
	output := stringsops.SinusoidalFrom(str)
	expected := "eHlol"

	if output != expected {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}

func TestSinusoidalFrom3(t *testing.T) {

	str := "data_structure"
	output := stringsops.SinusoidalFrom(str)
	expected := "ascedt_tutraru"

	if output != expected {
		t.Fatalf("Expected %v, received %v", expected, output)
	}
}
