// Ritual to start with Package main
package main

import "fmt"

// Board size
const size = 3

// Variable Empty board
var board [size][size]string

// STEP 1 = Create Board with empty space
func initBoard() {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			board[i][j] = " " // Fill with spaces
		}
	}
}

// STEP 2 - Print board
func printBoard() {
	fmt.Println("\n Welcome to Tic-Tac-Toe Game")
	for i, row := range board {
		fmt.Print(" ")
		for j, cell := range row {
			fmt.Print(cell)
			if j < size-1 {
				fmt.Print(" | ") // Column separators
			}
		}
		if i < size-1 {
			fmt.Println("\n-----------") // Row separators
		}
	}
	fmt.Println("\n")
}

// STEP 3 - Check Is board Full - This will check if any space is empty then ll return as False
func isBoardFull() bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == " " {
				return false
			}
		}
	}
	return true
}

// STEP 4 - Check if a player has won
func checkWinner(player string) bool {
	// Check rows and columns
	for i := 0; i < size; i++ {
		if (board[i][0] == player && board[i][1] == player && board[i][2] == player) || // Row
			(board[0][i] == player && board[1][i] == player && board[2][i] == player) { // Column
			return true
		}
	}
	// Check diagonals
	if (board[0][0] == player && board[1][1] == player && board[2][2] == player) || // Main diagonal
		(board[0][2] == player && board[1][1] == player && board[2][0] == player) { // Anti-diagonal
		return true
	}
	return false
}

// STEP 5 - Main game loop
func main() {
	initBoard()
	var player = "X"

	for {
		printBoard()
		fmt.Printf("Player %s, Please enter row and column (1-3, space-separated): ", player)

		var row, col int
		_, err := fmt.Scan(&row, &col)

		// Validate input
		if err != nil || row < 1 || row > 3 || col < 1 || col > 3 || board[row-1][col-1] != " " {
			fmt.Println("Invalid move Player! Try again between 1-2.")
			continue
		}

		// Place mark - Zero based index conversion
		board[row-1][col-1] = player

		// Check if player won - calling defined function above
		if checkWinner(player) {
			printBoard()
			fmt.Printf("🎉 Yooooooooo! Player %s wins!\n", player)
			break
		}

		// Check if it's a draw - Calling Draw func above
		if isBoardFull() {
			printBoard()
			fmt.Println("🤝 It's a draw!")
			break
		}

		// Switch turns - Else only one user keep playing
		if player == "X" {
			player = "O"
		} else {
			player = "X"
		}
	}
}
