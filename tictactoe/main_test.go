package main

import (
	"testing"
)

// Test initBoard function
func TestInitBoard(t *testing.T) {
	initBoard()
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] != " " {
				t.Errorf("Expected board[%d][%d] to be empty, but got '%s'", i, j, board[i][j])
			}
		}
	}
}

// Test isBoardFull function
func TestIsBoardFull(t *testing.T) {
	initBoard()
	if isBoardFull() {
		t.Errorf("Expected board to be not full, but got true")
	}

	// Fill the board
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			board[i][j] = "X"
		}
	}
	if !isBoardFull() {
		t.Errorf("Expected board to be full, but got false")
	}
}

// Test checkWinner function
func TestCheckWinner(t *testing.T) {
	initBoard()
	if checkWinner("X") {
		t.Errorf("Expected no winner initially, but got true")
	}

	// Test row win
	board[0][0], board[0][1], board[0][2] = "X", "X", "X"
	if !checkWinner("X") {
		t.Errorf("Expected X to win in row, but got false")
	}

	initBoard()
	// Test column win
	board[0][1], board[1][1], board[2][1] = "O", "O", "O"
	if !checkWinner("O") {
		t.Errorf("Expected O to win in column, but got false")
	}

	initBoard()
	// Test diagonal win
	board[0][0], board[1][1], board[2][2] = "X", "X", "X"
	if !checkWinner("X") {
		t.Errorf("Expected X to win diagonally, but got false")
	}
}
