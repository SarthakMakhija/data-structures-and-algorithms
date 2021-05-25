package stringsops_test

import (
	"reflect"
	"testing"

	stringsops "github.com/SarthakMakhija/data-structures-and-algorithms/strings_ops"
)

func TestStringDuplicatesGivenThereAreNoDuplicates(t *testing.T) {

	duplicates := stringsops.Duplicates("Sample")
	if len(duplicates) != 0 {
		t.Fatalf("Expecting no duplicates but found %v", duplicates)
	}
}

func TestStringDuplicatesGivenThereAreDuplicates_1(t *testing.T) {

	duplicates := stringsops.Duplicates("Hello")
	if !reflect.DeepEqual(duplicates, []string{"l"}) {
		t.Fatalf("Expecting a duplicate l but found %v", duplicates)
	}
}

func TestStringDuplicatesGivenThereAreDuplicates_2(t *testing.T) {

	duplicates := stringsops.Duplicates("Malayalam")
	if !reflect.DeepEqual(duplicates, []string{"l", "m", "a"}) {
		t.Fatalf("Expecting duplicates lma but found %v", duplicates)
	}
}
