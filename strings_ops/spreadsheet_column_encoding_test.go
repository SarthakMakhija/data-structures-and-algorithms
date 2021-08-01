package stringsops_test

import (
	stringsops "github.com/SarthakMakhija/data-structures-and-algorithms/strings_ops"
	"testing"
)

func TestEncodeExcelColumn1(t *testing.T) {
	column := "A"
	encoded := stringsops.EncodeExcelColumn(column)
	expected := 1

	if expected != encoded {
		t.Fatalf("Expected %v, received %v", expected, encoded)
	}
}

func TestEncodeExcelColumn2(t *testing.T) {
	column := "Z"
	encoded := stringsops.EncodeExcelColumn(column)
	expected := 26

	if expected != encoded {
		t.Fatalf("Expected %v, received %v", expected, encoded)
	}
}

func TestEncodeExcelColumn3(t *testing.T) {
	column := "AA"
	encoded := stringsops.EncodeExcelColumn(column)
	expected := 27

	if expected != encoded {
		t.Fatalf("Expected %v, received %v", expected, encoded)
	}
}

func TestEncodeExcelColumn4(t *testing.T) {
	column := "ZZ"
	encoded := stringsops.EncodeExcelColumn(column)
	expected := 702

	if expected != encoded {
		t.Fatalf("Expected %v, received %v", expected, encoded)
	}
}

func TestEncodeExcelColumn5(t *testing.T) {
	column := "CC"
	encoded := stringsops.EncodeExcelColumn(column)
	expected := 81

	if expected != encoded {
		t.Fatalf("Expected %v, received %v", expected, encoded)
	}
}

func TestEncodeExcelColumn6(t *testing.T) {
	column := "CCC"
	encoded := stringsops.EncodeExcelColumn(column)
	expected := 2109

	if expected != encoded {
		t.Fatalf("Expected %v, received %v", expected, encoded)
	}
}
