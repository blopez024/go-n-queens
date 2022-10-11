package main

import "fmt"

func main() {
	NQueens(8)
}

func NQueens(size int) {
	b := NewBoard(size)

	if Solution(b, 0, 0) {
		b.printBoard()
	} else {
		fmt.Println("No Solution")
	}

}

func Solution(b Board, x, y int) bool {
	q := Queen{x, y}

	// Successfully placed all Queens
	if x == len(b) {
		return true
	}

	if b.validPlacement(q) {
		b.placeQueen(x, q)
		Solution(b, x+1, 0)
	} else if y < len(b) {
		Solution(b, x, y+1)
	} else if x != 0 {
		prevY := b[x-1].y
		b.removeQueen(x - 1)
		Solution(b, x-1, prevY+1)
	}

	// Last Queen placed
	return b[len(b)-1].x != -1
}
