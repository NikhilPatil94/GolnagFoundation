package main

import (
	"testing"
)

// No Test is Required to initboard and print board as logic is not present

// testing Is Board Full
func TestIsBoardFull(t *testing.T) {
	initBoard() // empty board

	// Step 2: The board should not be full initially
	if isBoardFull() {
		t.Errorf("Expected board to be not full, but got true")
	}

	// Step 3: Fill the board completely
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			board[i][j] = "X" // Fill every cell
		}
	}

	// Step 4:board should be full
	if !isBoardFull() {
		t.Errorf("Expected board to be full, but got false")
	}
}

// Test for checking Winner Logic
func TestCheckWinner(t *testing.T) {
	// Test Winning Row
	board = [size][size]string{
		{"X", "X", "X"},
		{" ", "O", "O"},
		{" ", " ", " "},
	}
	if !checkWinner("X") {
		t.Errorf("Expected X to win with a row, but checkWinner returned false")
	}

	// Test Winning Column
	board = [size][size]string{
		{"O", "X", " "},
		{"O", "X", " "},
		{"O", " ", " "},
	}
	if !checkWinner("O") {
		t.Errorf("Expected O to win with a column, but checkWinner returned false")
	}

	// Test Winning Main Diagonal (\)
	board = [size][size]string{
		{"X", "O", "O"},
		{" ", "X", " "},
		{" ", " ", "X"},
	}
	if !checkWinner("X") {
		t.Errorf("Expected X to win with the main diagonal, but checkWinner returned false")
	}

	// Test Winning Anti-Diagonal (/)
	board = [size][size]string{
		{" ", "O", "X"},
		{" ", "X", "O"},
		{"X", " ", " "},
	}
	if !checkWinner("X") {
		t.Errorf("Expected X to win with the anti-diagonal, but checkWinner returned false")
	}

	// Test No Winner (Game Still Ongoing)
	board = [size][size]string{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"O", "X", " "},
	}
	if checkWinner("X") || checkWinner("O") {
		t.Errorf("Expected no winner, but checkWinner returned true")
	}
}

// Test to valid imputs given by players
func TestIsValidMove(t *testing.T) {
	// Setup: Empty board
	initBoard()

	// Test Case 1: Valid move in an empty cell
	if !isValidMove(1, 1) {
		t.Errorf("Expected true for an empty cell, but got false")
	}

	// Test Case 2: Out-of-bounds moves
	if isValidMove(0, 1) {
		t.Errorf("Expected false for row=0 (out of bounds), but got true")
	}
	if isValidMove(4, 2) {
		t.Errorf("Expected false for row=4 (out of bounds), but got true")
	}

	// Test Case 3: Move on an occupied cell
	board[0][0] = "X" // Mark (1,1) as occupied
	if isValidMove(1, 1) {
		t.Errorf("Expected false for occupied cell, but got true")
	}
}
