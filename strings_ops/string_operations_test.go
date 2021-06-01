package stringsops_test

import (
	"reflect"
	"testing"

	stringsops "github.com/SarthakMakhija/data-structures-and-algorithms/strings_ops"
)

func TestStringDuplicatesGivenThereAreNoDuplicates_1(t *testing.T) {

	duplicates := stringsops.Duplicates1("Sample")
	if len(duplicates) != 0 {
		t.Fatalf("Expecting no duplicates but found %v", duplicates)
	}
}

func TestStringDuplicatesGivenThereAreDuplicates_1(t *testing.T) {

	duplicates := stringsops.Duplicates1("Hello")
	if !reflect.DeepEqual(duplicates, []string{"l"}) {
		t.Fatalf("Expecting a duplicate l but found %v", duplicates)
	}
}

func TestStringDuplicatesGivenThereAreDuplicates_2(t *testing.T) {

	duplicates := stringsops.Duplicates1("Malayalam")
	if !reflect.DeepEqual(duplicates, []string{"l", "m", "a"}) {
		t.Fatalf("Expecting duplicates lma but found %v", duplicates)
	}
}

func TestStringDuplicatesGivenThereAreNoDuplicates_2(t *testing.T) {

	duplicates := stringsops.Duplicates2("sample")
	if len(duplicates) != 0 {
		t.Fatalf("Expecting no duplicates but found %v", duplicates)
	}
}

func TestStringDuplicatesGivenThereAreNoDuplicates_3(t *testing.T) {

	duplicates := stringsops.Duplicates2("Sample")
	if len(duplicates) != 0 {
		t.Fatalf("Expecting no duplicates but found %v", duplicates)
	}
}

func TestStringDuplicatesGivenThereAreDuplicates_3(t *testing.T) {

	duplicates := stringsops.Duplicates2("Hello")
	if !reflect.DeepEqual(duplicates, []string{"l"}) {
		t.Fatalf("Expecting a duplicate l but found %v", duplicates)
	}
}

func TestStringDuplicatesGivenThereAreDuplicates_4(t *testing.T) {

	duplicates := stringsops.Duplicates1("Malayalam")
	if !reflect.DeepEqual(duplicates, []string{"l", "m", "a"}) {
		t.Fatalf("Expecting duplicates lma but found %v", duplicates)
	}
}

func TestStringDuplicatesGivenThereAreNoDuplicates_4(t *testing.T) {

	duplicateExists := stringsops.Duplicates3("sample")
	if duplicateExists != false {
		t.Fatalf("Expecting no duplicates but found %v", duplicateExists)
	}
}

func TestStringDuplicatesGivenThereAreNoDuplicates_5(t *testing.T) {

	duplicateExists := stringsops.Duplicates3("Sample")
	if duplicateExists != false {
		t.Fatalf("Expecting no duplicates but found %v", duplicateExists)
	}
}

func TestStringDuplicatesGivenThereAreNoDuplicates_6(t *testing.T) {

	duplicateExists := stringsops.Duplicates3("Hello")
	if duplicateExists != true {
		t.Fatalf("Expecting duplicates but found %v", duplicateExists)
	}
}

func TestStringDuplicatesGivenThereAreNoDuplicates_7(t *testing.T) {

	duplicateExists := stringsops.Duplicates3("Finding")
	if duplicateExists != true {
		t.Fatalf("Expecting duplicates but found %v", duplicateExists)
	}
}

func TestStringDuplicatesGivenThereAreNoDuplicates_8(t *testing.T) {

	duplicateExists := stringsops.Duplicates3("Malayalam")
	if duplicateExists != true {
		t.Fatalf("Expecting duplicates but found %v", duplicateExists)
	}
}
