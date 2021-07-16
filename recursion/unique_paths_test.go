package recursion

import "testing"

func TestPathsFromUpperLeftmostToLowerRightmostIn1(t *testing.T) {
	totalsPaths := TotalsPathsFromUpperLeftmostToLowerRightmostIn(3, 3)
	if totalsPaths != 6 {
		t.Fatalf("Expected totalsPaths to be 6, rececied %v", totalsPaths)
	}
}

func TestPathsFromUpperLeftmostToLowerRightmostIn2(t *testing.T) {
	totalsPaths := TotalsPathsFromUpperLeftmostToLowerRightmostIn(3, 7)
	if totalsPaths != 28 {
		t.Fatalf("Expected totalsPaths to be 28, rececied %v", totalsPaths)
	}
}
