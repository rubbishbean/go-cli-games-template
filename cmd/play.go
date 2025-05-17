package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-cli-games-template/games/guessnumber"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play a game",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1. Guess the Number")
		fmt.Println("2. Tic Tac Toe")
		// more...
		guessnumber.Play()
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}