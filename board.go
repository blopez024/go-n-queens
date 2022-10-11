package main

import (
	"fmt"
	"math"
)

type Queen struct {
	x int
	y int
}

type Board []Queen

func NewBoard(s int) Board {
	b := make(Board, s)
	for i := 0; i < s; i++ {
		b[i] = Queen{x: -1, y: -1}
	}
	return b
}

func AttacksQueen(q1, q2 Queen) bool {
	// Horizontal
	if q1.x == q2.x {
		return true
	}
	// Vertical
	if q1.y == q2.y {
		return true
	}
	// Diagonal
	if math.Abs(float64(q1.x-q2.x)) == math.Abs(float64(q1.y-q2.y)) {
		return true
	}

	return false
}

func (b Board) validPlacement(q Queen) bool {
	// Out of Bounds
	if q.x >= len(b) || q.y >= len(b) {
		return false
	}

	// Previous queens placement
	for qx := q.x - 1; qx > -1; qx-- {
		q2 := b[qx]
		if AttacksQueen(q, q2) {
			return false
		}
	}
	return true
}

func (b Board) placeQueen(x int, q Queen) {
	b[x] = q
}

func (b Board) removeQueen(x int) {
	b[x] = Queen{x: -1, y: -1}
}

func (b Board) printBoard() {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b); j++ {
			if j != b[i].y {
				fmt.Print("- ")
			} else {
				fmt.Print("♕ ")
			}
		}
		fmt.Println()
	}
}
