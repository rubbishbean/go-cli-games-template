package tictactoe

import "fmt"

// Game struct holds the state of the Tic Tac Toe game
type Game struct {
	Board [3][3]rune
	Turn  rune // 'X' or 'O'
}

// Play runs the Tic Tac Toe game
func Play() {
	fmt.Println("[TODO] Implement Tic Tac Toe game logic here.")
	// Suggested steps:
	// - Initialize game state (use NewGame())
	// - Alternate turns between X and O
	// - Check for win/draw conditions after each move
	// - Print the board
}

// NewGame creates and returns a new game instance
func NewGame() *Game {
	var board [3][3]rune
	for i := range board {
		for j := range board[i] {
			board[i][j] = ' '
		}
	}
	return &Game{
		Board: board,
		Turn:  'X',
	}
}
