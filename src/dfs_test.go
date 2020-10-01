package main

import (
	"testing"
)
// Test can't be test
func TestPuzzleEqualTrue(t *testing.T) {
	 a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}
	 b := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}

	 if(puzzleEqual(a,b) != true){
		t.Error("should return true")
	 }
}

func TestPuzzleEqualFalse(t *testing.T) {
	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 1}}
	b := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}

	if(puzzleEqual(a,b) != true){
	   t.Error("should return false")
	}
}
