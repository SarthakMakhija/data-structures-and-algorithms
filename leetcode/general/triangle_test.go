package general

import "testing"

func TestMinimumTriangleSum1(t *testing.T) {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	minimumSum := MinimumTriangleSum(triangle)

	if minimumSum != 11 {
		t.Fatalf("Expected 11, received %v", minimumSum)
	}
}

func TestMinimumTriangleSum2(t *testing.T) {
	triangle := [][]int{
		{2},
		{4, 3},
		{7, 6, 5},
		{8, 1, 4, 3},
	}
	minimumSum := MinimumTriangleSum(triangle)

	if minimumSum != 12 {
		t.Fatalf("Expected 12, received %v", minimumSum)
	}
}

func TestMinimumTriangleSum3(t *testing.T) {
	triangle := [][]int{
		{2},
		{4, 5},
		{6, 7, 8},
		{9, 10, 11, 12},
	}
	minimumSum := MinimumTriangleSum(triangle)

	if minimumSum != 21 {
		t.Fatalf("Expected 21, received %v", minimumSum)
	}
}

func TestMinimumTriangleSum4(t *testing.T) {
	triangle := [][]int{
		{2},
		{5, 4},
		{8, 6, 7},
		{10, 11, 12, 9},
	}
	minimumSum := MinimumTriangleSum(triangle)

	if minimumSum != 22 {
		t.Fatalf("Expected 22, received %v", minimumSum)
	}
}
