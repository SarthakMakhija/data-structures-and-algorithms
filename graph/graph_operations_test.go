package graph

import (
	"reflect"
	"testing"
)

/**
Graph is represented as a 2D array
Row1 and Row2 indicate the source and destination vertices
Row3 indicate the weight from vertex1 to vertex2
*/
func TestShortestPathFrom1(t *testing.T) {
	directedGraph := [3][]int{
		{1, 2, 1, 2, 3, 4, 5, 5},
		{2, 3, 3, 4, 5, 6, 4, 6},
		{2, 1, 4, 7, 3, 1, 2, 5},
	}
	expectedPath := []int{1, 2, 3, 5, 4, 6}
	shortestPath := ShortestPathFrom(1, 6, directedGraph)
	if shortestPath.weight != 9 {
		t.Fatalf("Expected weight %v, received %v", 9, shortestPath.weight)
	}
	if !reflect.DeepEqual(shortestPath.path, expectedPath) {
		t.Fatalf("Expected path %v, received %v", expectedPath, shortestPath.path)
	}
}

func TestShortestPathFrom2(t *testing.T) {
	directedGraph := [3][]int{
		{1, 2, 1, 2, 3, 4, 5, 5},
		{2, 3, 3, 4, 5, 6, 4, 6},
		{2, 1, 4, 7, 3, 1, 2, 5},
	}
	expectedPath := []int{3, 5, 4, 6}
	shortestPath := ShortestPathFrom(3, 6, directedGraph)
	if shortestPath.weight != 6 {
		t.Fatalf("Expected weight %v, received %v", 6, shortestPath.weight)
	}
	if !reflect.DeepEqual(shortestPath.path, expectedPath) {
		t.Fatalf("Expected path %v, received %v", expectedPath, shortestPath.path)
	}
}
