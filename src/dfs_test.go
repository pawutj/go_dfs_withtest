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
	s1 := State{id: 0}
	s2 := State{id: 1}
	q := Queue{data: []State{s1, s2}}
	p1 := q.pop()
	p2 := q.pop()
	if p1.id != 0 {
		t.Error("should return 0")
	}
	if p2.id != 1 {
		t.Error("should return 1")
	}
}

func TestQueuePush(t *testing.T) {
	q := Queue{data: []State{}}
	s1 := State{id: 0}
	q.push(s_1)
	p1 := q.pop()
	if p1.id != 0 {
		t.Error("shold return 0")
	}

}
