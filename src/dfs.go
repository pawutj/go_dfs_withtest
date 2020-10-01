package main

import "fmt"

var Goal = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}
var Init = [][]int{{4, 1, 2}, {7, 0, 3}, {8, 5, 6}}
var b Queue


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
	puzzle   [][]int
	oldState []string
}

type Queue struct {
	data []interface{}
}

func (q Queue) pop() interface{} {
	head := q.data[0]
	q.data = q.data[1:]
	return head
}

func (q Queue) push(a interface{}) {
	q.data = append(q.data, a)
}

func bfs() {

}

func main() {

	c := copyPuzzle(Goal)
	fmt.Println(findZeroLocation(Goal))
	d := swapElement(c, 0, 0, 1, 1)

	fmt.Println(Goal)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(puzzleEqual(Goal, c))

	//a:=[]interface{}{1,2,3}
	b = Queue{data: []interface{}{1, 2, 3}}
	fmt.Println(b.pop())
}
