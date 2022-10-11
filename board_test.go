package main

import (
	"testing"
)

func TestAttack(t *testing.T) {
	q1 := Queen{x: 1, y: 1}
	q2 := Queen{x: 1, y: 8}

	if AttacksQueen(q1, q2) != true {
		t.Errorf("Expected Q1: (%d,%d) to attack Q2: (%d,%d)", q1.x, q1.y, q2.x, q2.y)
	}
	if AttacksQueen(q2, q1) != true {
		t.Errorf("Expected Q1: (%d,%d) to attack Q2: (%d,%d)", q2.x, q2.y, q1.x, q1.y)
	}

	q1 = Queen{x: 1, y: 8}
	q2 = Queen{x: 8, y: 8}

	if AttacksQueen(q1, q2) != true {
		t.Errorf("Expected Q1: (%d,%d) to attack Q2: (%d,%d)", q1.x, q1.y, q2.x, q2.y)
	}
	if AttacksQueen(q2, q1) != true {
		t.Errorf("Expected Q1: (%d,%d) to attack Q2: (%d,%d)", q2.x, q2.y, q1.x, q1.y)
	}

	q1 = Queen{x: 1, y: 1}
	q2 = Queen{x: 8, y: 8}

	if AttacksQueen(q1, q2) != true {
		t.Errorf("Expected Q1: (%d,%d) to attack Q2: (%d,%d)", q1.x, q1.y, q2.x, q2.y)
	}
	if AttacksQueen(q2, q1) != true {
		t.Errorf("Expected Q1: (%d,%d) to attack Q2: (%d,%d)", q2.x, q2.y, q1.x, q1.y)
	}
}

func TestPlacingQueens(t *testing.T) {
	board := NewBoard(4)

	q1 := Queen{x: 0, y: 1}
	q2 := Queen{x: 1, y: 3}
	q3 := Queen{x: 2, y: 0}
	q4 := Queen{x: 3, y: 2}

	board.placeQueen(0, q1)
	board.placeQueen(1, q2)
	board.placeQueen(2, q3)
	board.placeQueen(3, q4)

	if board[0].x != 0 || board[0].y != 1 {
		t.Errorf("Expected Q1 to be at (%d,%d) not (%d,%d)", 0, 1, q1.x, q1.y)
	}
	if board[1].x != 1 || board[1].y != 3 {
		t.Errorf("Expected Q2 to be at (%d,%d) not (%d,%d)", 1, 3, q2.x, q2.y)
	}
	if board[2].x != 2 || board[2].y != 0 {
		t.Errorf("Expected Q3 to be at (%d,%d) not (%d,%d)", 2, 0, q3.x, q3.y)
	}
	if board[3].x != 3 || board[3].y != 2 {
		t.Errorf("Expected Q4 to be at (%d,%d) not (%d,%d)", 3, 2, q4.x, q4.y)
	}
}
