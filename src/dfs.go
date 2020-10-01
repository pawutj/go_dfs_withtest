package main

import "fmt"

var Goal = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}
var Init = [][]int{{4, 1, 2}, {7, 0, 3}, {8, 5, 6}}
var bfsArray Queue

func findZeroLocation(puzzle [][]int) (int, int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if puzzle[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func copyPuzzle(puzzle [][]int) [][]int {
	_puzzle := make([][]int, 3)
	for i := 0; i < 3; i++ {
		_puzzle[i] = make([]int, 3)
		copy(_puzzle[i], puzzle[i])
	}
	return _puzzle
}

func swapElement(puzzle [][]int, x1 int, y1 int, x2 int, y2 int) [][]int {

	_puzzle := copyPuzzle(puzzle)

	temp := _puzzle[x1][y1]
	_puzzle[x1][y1] = _puzzle[x2][y2]
	_puzzle[x2][y2] = temp
	return _puzzle
}

func puzzleEqual(a [][]int, b [][]int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}

	}
	return true
}

type State struct {
	id       int
	puzzle   [][]int
	oldState []string
}

type Queue struct {
	data []State
}

func (q *Queue) pop() State {
	head := q.data[0]
	q.data = q.data[1:]
	return head
}

func (q *Queue) push(a State) {
	q.data = append(q.data, a)
}

func bfs() bool {
	_c := bfsArray.pop()
	c := _c.puzzle
	//fmt.Println(c)
	if puzzleEqual(c, Goal) {
		fmt.Println("END")
		fmt.Println(_c.oldState)
		return true
	}
	x, y := findZeroLocation(c)
	if x < 2 {
		NewPuzzle := swapElement(c, x, y, x+1, y)
		NewState := append(_c.oldState, "D")
		newState := State{puzzle: NewPuzzle, oldState: NewState}
		bfsArray.push(newState)
	}

	if x > 0 {
		NewPuzzle := swapElement(c, x, y, x-1, y)
		NewState := append(_c.oldState, "U")
		newState := State{puzzle: NewPuzzle, oldState: NewState}
		bfsArray.push(newState)
	}

	if y > 0 {
		NewPuzzle := swapElement(c, x, y, x, y-1)
		NewState := append(_c.oldState, "L")
		newState := State{puzzle: NewPuzzle, oldState: NewState}
		bfsArray.push(newState)
	}

	if y < 2 {
		NewPuzzle := swapElement(c, x, y, x, y+1)
		NewState := append(_c.oldState, "R")
		newState := State{puzzle: NewPuzzle, oldState: NewState}
		bfsArray.push(newState)
	}

	return false
}

func main() {

	bfsArray.push(State{puzzle: Init, oldState: []string{}})

	for end := false; end != true; {
		end = bfs()
	}

}
