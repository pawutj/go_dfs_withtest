package main

import (
	"testing"
)

// Test can't be test
func TestPuzzleEqualTrue(t *testing.T) {
	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}
	b := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}

	if puzzleEqual(a, b) != true {
		t.Error("should return true")
	}
}

func TestPuzzleEqualFalse(t *testing.T) {
	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 1}}
	b := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}

	if puzzleEqual(a, b) != false {
		t.Error("should return false")
	}
}

func TestFindZeroLocation(t *testing.T) {
	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}
	x, y := findZeroLocation(a)
	if x != 2 || y != 2 {
		t.Error("should return 2,2")
	}

}

func TestQueuePop(t *testing.T) {

	q := Queue{data: []interface{}{1, 2, 3}}
	c_1 := q.pop()
	c_2 := q.pop()
	if c_1 != 1 {
		t.Error("should return 1")
	}
	if c_2 != 2 {
		t.Error("should return 2")
	}
}

func TestQueuePush(t *testing.T){
	q := Queue{data : []interface{}{}}
	q.push(1)
	c := q.pop()
	if(c!=1){
		t.Error("shold return 1")
	}

}